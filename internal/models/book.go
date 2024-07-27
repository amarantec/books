package models

type Book struct {
	Id		       	int64		`json:"id"`
	Title 		    	string 		`json:"title"`
	Description 		string 		`json:"description"`
	Genre		      	[]string	`json:"genre"`
	Author		    	[]string 	`json:"author"`
	ImageURL      		string    	`json:"image_url"`
	CategoryId	  	int64 		`json:"category_id"`
	UserId		    	int64 		`json:"user_id"`
}
