// Package service
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 文章功能
 * @File:  article
 * @Version: 1.0.0
 * @Date: 2022/7/4 20:20
 */
package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"r0Website-server/global"
	"r0Website-server/models"
	"r0Website-server/models/views"
	"r0Website-server/utils"
	"time"
)

type ArticleService struct{}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

const ArticleColl = "articles"

// ArticleADDFile 通过上传文件增加文章
func (article *ArticleService) ArticleADDFile(
	params views.AdminArticleAddFileVo, id string,
) (vo *views.AdminArticleAddFileResultVo, err error) {
	var result views.AdminArticleAddFileResultVo
	var input models.Article
	// 获取文章内容
	open, err := params.File.Open()
	if err != nil {
		global.Logger.Error(err)
		return nil, err
	}
	content, err := ioutil.ReadAll(open)
	if err != nil {
		global.Logger.Error(err)
		return nil, errors.New("ArticleADDFile: 获取文件内容失败")
	}
	input.Markdown = string(content)
	UpdateArticleMetaByParams(&input, params, id)
	var insertResult *mongo.InsertOneResult
	insertResult, err = global.DBEngine.Collection(ArticleColl).InsertOne(context.TODO(), input)
	if err != nil {
		global.Logger.Error(err)
		return nil, &models.UniqueError{UniqueField: "article->_id", Msg: input.Id.Hex(), Count: 1}
	}
	result = views.AdminArticleAddFileResultVo{Title: input.Title, Id: insertResult.InsertedID.(primitive.ObjectID)}
	return &result, err
}

// ArticleADDForm 通过提交表单增加文章
func (article *ArticleService) ArticleADDForm(
	params views.AdminArticleAddFormVo, id string,
) (vo *views.AdminArticleAddFormResultVo, err error) {
	var result views.AdminArticleAddFormResultVo
	var input models.Article
	/*if id, err = checkAndPatchArticleUuid(id); err != nil {
		return nil, err
	}*/
	input.Markdown = params.Markdown
	UpdateArticleMetaByParams(&input, params, id)
	var insertResult *mongo.InsertOneResult
	insertResult, err = global.DBEngine.Collection(ArticleColl).InsertOne(context.TODO(), input)
	if err != nil {
		global.Logger.Error(err)
		return nil, &models.UniqueError{UniqueField: "article->_id", Msg: input.Id.Hex(), Count: 1}
	}
	result = views.AdminArticleAddFormResultVo{Title: input.Title, Id: insertResult.InsertedID.(primitive.ObjectID)}
	return &result, err
}

// ArticleBaseSearch 基础权限的文章搜索功能
// 可选作者，可选模糊内容，可选两种时间排序，可选id
// **如果使用id检索，其他检索全部失效**
// 分页 利用opt构造的skip和limit
// FOLLOWS:
// - https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/text/
// - https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/limit/
// db.articles.find({$text:{$search:"xxx"}, author:"xxx"},{score:{$meta : "textScore"}})
//	.sort({"update_time":-1, score:{$meta : "textScore"}})
func (article *ArticleService) ArticleBaseSearch(
	params views.BaseArticleSearchVo, id string,
) (vo *views.BaseArticleSearchResultVo, err error) {
	var result views.BaseArticleSearchResultVo
	searchText := params.SearchText
	author := params.Author
	pageNumber := params.PageNumber
	pageSize := params.PageSize
	pageSkip := int64(0)
	// 是一个补丁，防止出现WriteNull错误，丑陋的解决方法
	filter := bson.D{}
	sort := bson.D{}
	projection := bson.D{}
	opts := options.Find()
	result.Articles = []views.SingleBaseArticleSearchResultVo{}
	// 构造排序bson
	if params.UpdateTimeSort.SortFlag == true && params.CreateTimeSort.SortFlag == true {
		return nil, errors.New("ArticleBaseSearch: " + "不能同时指定UpdateTime和CreateTime的排序")
	}
	if params.UpdateTimeSort.SortFlag == true {
		sort = bson.D{{"update_time", params.UpdateTimeSort.SortDirection}}
		opts = opts.SetSort(sort)
	}
	if params.CreateTimeSort.SortFlag == true {
		sort = bson.D{{"create_time", params.CreateTimeSort.SortDirection}}
		opts = opts.SetSort(sort)
	}
	// 构造搜索/过滤bson
	if author != "" {
		filter = append(filter, bson.E{Key: "author", Value: author})
	}
	// 如果使用id检索，其他检索全部失效
	if id != "" {
		if bsonId, err := primitive.ObjectIDFromHex(utils.String2HexString24(id)); err != nil {
			global.Logger.Error(err)
		} else {
			filter = bson.D{{"_id", bsonId}}
			result.Msg = "使用id检索，其他检索条件全部失效"
		}
	}
	// 如果使用id检索，其他检索全部失效，并不应该在排序段增加score
	if searchText != "" && id == "" {
		// 如果包含模糊搜素，那么需要在sort段和投影段增加条件
		filter = append(filter, bson.E{Key: "$text", Value: bson.D{{"$search", searchText}}})
		sort = append(sort, bson.E{Key: "score", Value: bson.D{{"$meta", "textScore"}}})
		// 构造投影bson
		projection = append(projection, bson.E{Key: "score", Value: bson.D{{"$meta", "textScore"}}})
		// 这里会引起opts的二次SetSort，但目前还没发现问题
		opts = opts.SetSort(sort).SetProjection(projection)
	}
	// 构造分页, 页码从1开始，需要同时指定才能生效
	if pageNumber != 0 && pageSize != 0 {
		pageSkip = (pageNumber - 1) * pageSize
		opts = opts.SetLimit(pageSize).SetSkip(pageSkip)
	}
	global.Logger.Infof("ArticleBaseSearch -> Mongo: \n\t[ %+v | %+v | %+v ]", filter, sort, projection)
	cursor, err := global.DBEngine.Collection(ArticleColl).Find(context.TODO(), filter, opts)
	if err != nil {
		global.Logger.Error(err)
	}
	if err = cursor.All(context.TODO(), &result.Articles); err != nil {
		global.Logger.Error(err)
	}
	for index, val := range result.Articles {
		result.Articles[index].UpdateTime = val.UpdateTime.Local()
		result.Articles[index].CreateTime = val.CreateTime.Local()
	}
	// defer 关闭游标
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			global.Logger.Error(err)
		}
	}(cursor, context.TODO())
	return &result, nil
}

// UpdateArticleMetaByParams 用输入的参数更新文章元信息
func UpdateArticleMetaByParams(input *models.Article, params interface{}, uuid string) {
	switch params.(type) {
	case views.AdminArticleAddFormVo, views.AdminArticleAddFileVo:
		var meta views.AdminArticleAddMetaVo
		if value, ok := params.(views.AdminArticleAddFormVo); ok {
			meta = value.AdminArticleAddMetaVo
		} else if value, ok := params.(views.AdminArticleAddFileVo); ok {
			meta = value.AdminArticleAddMetaVo
		}
		// input.Uuid = uuid
		var err error
		uuid = utils.String2HexString24(uuid)
		if input.Id, err = primitive.ObjectIDFromHex(uuid); err != nil {
			global.Logger.Error(err)
		}
		input.Title = meta.Title
		input.Author = meta.Author
		input.Synopsis = meta.Synopsis
		input.DeleteFlag = false
		input.DraftFlag = meta.DraftFlag
		// input.Detail = ""
		// input.Markdown = meta.Markdown
		input.Overhead = meta.Overhead
		input.ArtLength = 0
		input.ReadsNumber = 0
		input.CommentsNumber = 0
		input.PraiseNumber = 0
		input.Tags = meta.Tags
		input.Categories = meta.Categories
		var curTime = time.Now()
		input.UpdateTime = curTime
		input.CreateTime = curTime
		// 接下来“修补”模型的值
		// input.Detail = utils.Markdown2Html(input.Markdown)
		// input.ArtLength = len(input.Markdown)
		wordCounter := utils.WordCounter{}
		wordCounter.Stat(input.Markdown)
		input.ArtLength = int64(wordCounter.Total)
		input.MdWords = utils.WordSplitForSearching(input.Markdown)
		input.TitleWords = utils.WordSplitForSearching(input.Title)
	}

}

// checkAndPatchArticleUuid 检查并修补文章uuid的值
func checkAndPatchArticleUuid(uuid string) (string, error) {
	var assignedUuid = true
	if uuid == "" {
		// 空id 将自动生成uuid
		uuid = utils.GenSonyflake()
		assignedUuid = false
	}
	if articleUuidCount := articleCountByUUid(uuid); articleUuidCount > 0 {
		// 表示id重复 重复生成一次 两次Sonyflake碰撞的概率非常低
		// 但对于指定了idq的POST请求的碰撞 应该直接给UniqueError
		if assignedUuid {
			return "", &models.UniqueError{UniqueField: "article-uuid", Msg: uuid, Count: articleUuidCount}
		} else {
			uuid = utils.GenSonyflake()
		}
	}
	return uuid, nil
}

// articleCountByUUid 统计有多少文章有此uuid
func articleCountByUUid(uuid string) int64 {
	filter := bson.M{"uuid": uuid}
	count, err := global.DBEngine.Collection(ArticleColl).CountDocuments(context.TODO(), filter)
	if err != nil {
		global.Logger.Error(err)
	}
	return count
}
