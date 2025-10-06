package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       120,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/e/e3/Oranges_-_whole-halved-segment.jpg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is Green",
		Price:       120,
		ImgUrl:      "https://static.vecteezy.com/system/resources/thumbnails/012/086/172/small_2x/green-apple-with-green-leaf-isolated-on-white-background-vector.jpg",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is Yellow",
		Price:       60,
		ImgUrl:      "https://www.bobtailfruit.co.uk/media/mageplaza/blog/post/4/2/42e9as7nataai4a6jcufwg.jpeg",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Watermelon",
		Description: "Watermelon is Sweet",
		Price:       300,
		ImgUrl:      "https://media.post.rvohealth.io/wp-content/uploads/sites/3/2025/01/benefits-watermelon-GettyImages-1331129269-Thumb.jpg",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "MANGO is my favorute ",
		Price:       300,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
}
