package main

import (
    "context"
    "fmt"
    "net/http"

	"github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "serverdevops/models"
)

var client *mongo.Client

func connectToMongo() {
    var err error
    client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://mongo-db:27017"))
    if err != nil {
        panic(err)
    }

    err = client.Connect(context.TODO())
    if err != nil {
        panic(err)
    }

    fmt.Println("Connected to MongoDB!")
}

func main() {
    connectToMongo()
    router := gin.Default()

	// Thêm middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Cho phép miền này
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))


    // CRUD routes
    router.GET("api/v1/books", getBooks)
    router.GET("api/v1/books/:id", getBook)
    router.POST("api/v1/books", createBook)
    router.PUT("api/v1/books/:id", updateBook)
    router.DELETE("api/v1/books/:id", deleteBook)

    router.Run(":6080")
}

func getBooks(c *gin.Context) {
    collection := client.Database("bookstore").Collection("books")
    cur, err := collection.Find(context.TODO(), bson.D{{}})
    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    var books []models.Book
    if err = cur.All(context.TODO(), &books); err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)
    collection := client.Database("bookstore").Collection("books")
    var book models.Book
    err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&book)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, err)
        return
    }
    collection := client.Database("bookstore").Collection("books")
    book.ID = primitive.NewObjectID()
    _, err := collection.InsertOne(context.TODO(), book)
    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusCreated, book)
}

func updateBook(c *gin.Context) {
    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, err)
        return
    }
    collection := client.Database("bookstore").Collection("books")
    _, err := collection.UpdateOne(
        context.TODO(),
        bson.M{"_id": objID},
        bson.D{
            {"$set", bson.D{
                {"title", book.Title},
                {"author", book.Author},
                {"year", book.Year},
            }},
        },
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusOK, book)
}

func deleteBook(c *gin.Context) {
    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)
    collection := client.Database("bookstore").Collection("books")
    _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusNoContent, nil)
}