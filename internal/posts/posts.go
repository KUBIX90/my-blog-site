package posts

type Post struct {
	Title   string
	Content string
}

func GetPosts() []*Post {
	return []*Post{
		{Title: "Post 1", Content: "Content of post 1"},
		{Title: "Post 2", Content: "Content of post 2"},
	}
}
