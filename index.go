package main

import (
	"context"
	"fmt"
	"log"
	prisma "workspace/prisma-gin/prisma-client"

	"github.com/gin-gonic/gin"
)

func main() {

	client := prisma.New(nil)

	r := gin.Default()
	ctx := context.Background()

	r.GET("/posts/published", func(c *gin.Context) {
		published := true
		posts, err := client.Posts(&prisma.PostsParams{
			Where: &prisma.PostWhereInput{
				Published: &published,
			},
		},
		).Exec(ctx)

		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{
			"posts": posts,
		})
	})

	r.GET("/post/:id", func(c *gin.Context) {
		id := c.Param("id")

		post, err := client.Post(prisma.PostWhereUniqueInput{
			ID: &id,
		},
		).Exec(ctx)

		if err != nil {
			log.Printf("%v", err)
		}

		c.JSON(200, gin.H{
			"post": post,
		})
	})

	r.GET("/posts/user/:userID", func(c *gin.Context) {
		userID := c.Param("userID")

		posts, err := client.Posts(&prisma.PostsParams{
			Where: &prisma.PostWhereInput{
				Author: &prisma.UserWhereInput{
					ID: &userID,
				},
			},
		},
		).Exec(ctx)

		if err != nil {
			log.Printf("%v", err)
		}
		c.JSON(200, gin.H{
			"posts": posts,
		})
	})

	r.POST("/post/draft", func(c *gin.Context) {
		var p map[string]string
		c.BindJSON(&p)

		title := p["title"]
		userID := p["userId"]

		post, err := client.CreatePost(prisma.PostCreateInput{
			Title: title,
			Author: &prisma.UserCreateOneWithoutPostsInput{
				Connect: &prisma.UserWhereUniqueInput{
					ID: &userID,
				},
			},
		},
		).Exec(ctx)

		if err != nil {
			log.Printf("%v", err)
		}
		c.JSON(200, gin.H{
			"post": post,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		var u map[string]string
		c.BindJSON(&u)

		name := u["name"]

		user, err := client.CreateUser(prisma.UserCreateInput{
			Name: name,
		},
		).Exec(ctx)

		if err != nil {
			log.Printf("%v", err)
		}
		c.JSON(200, gin.H{
			"user": user,
		})
	})

	r.PUT("/post/publish/:postID", func(c *gin.Context) {
		postID := c.Param("postID")
		published := true
		post, err := client.UpdatePost(prisma.PostUpdateParams{
			Where: prisma.PostWhereUniqueInput{
				ID: &postID,
			},
			Data: prisma.PostUpdateInput{
				Published: &published,
			},
		},
		).Exec(ctx)

		if err != nil {
			log.Printf("%v", err)
		}

		c.JSON(200, gin.H{
			"post": post,
		})
	})

	fmt.Println("Server is running on http://localhost:8080")
	r.Run()
}
