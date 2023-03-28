package entity

type JSend struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

const (
	ConstJSendSuccess = "success"
	ConstJSendFail    = "fail"
)

/*
EXAMPLE:
{
    status : "success",
    data : {
        "posts" : [
            { "id" : 1, "title" : "A blog post", "body" : "Some useful content" },
            { "id" : 2, "title" : "Another blog post", "body" : "More content" },
        ]
     }
}
*/
