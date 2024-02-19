package storage

import (
	"database/sql"
	"fmt"
	"github.com/zelinskayas/GoBasic/goodRestApiWithDBAuth/internal/app/models"
	"log"
)

// instance of article repository (model inteface)
type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "dbo.articles"
)

// добавим статью в бд
func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("USE MyLocalDB INSERT INTO %s (title, author, content) VALUES(@val1,@val2,@val3) SELECT SCOPE_IDENTITY()", tableArticle)
	if err := ar.storage.db.QueryRow(query, sql.Named("val1", a.Title), sql.Named("val2", a.Author), sql.Named("val3", a.Content)).Scan(&a.ID); err != nil {
		log.Printf(query)
		return nil, err
	}
	return a, nil
}

// удаляем статью по id
func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {
	article, ok, err := ar.FindArticleById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("USE MyLocalDB DELETE FROM %s WHERE id = @id", tableArticle)
		_, err := ar.storage.db.Exec(query, sql.Named("id", id))
		if err != nil {
			return nil, err
		}
	}
	return article, nil
}

// получаем статью по id
func (ar *ArticleRepository) FindArticleById(id int) (*models.Article, bool, error) {
	articles, err := ar.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var articleFinded *models.Article
	for _, u := range articles {
		if u.ID == id {
			articleFinded = u
			founded = true
			break
		}
	}
	return articleFinded, founded, nil
}

// получим все статьи в бд
func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("USE MyLocalDB SELECT id, title, author, content FROM %s", tableArticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//подготовим куда будем читать
	articles := make([]*models.Article, 0)
	for rows.Next() {
		a := models.Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	return articles, nil
}
