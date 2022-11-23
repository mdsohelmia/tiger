package core

// type Pagination struct {
// 	PerPage    int   `json:"per_page"`
// 	Count      int   `json:"count"`
// 	Page       int   `json:"page"`
// 	TotalCount int64 `json:"total_count"`
// 	TotalPages int64 `json:"total_pages"` //Total Page
// }

// func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
// 	perPage := 2
// 	page := 1
// 	sort := "created_at asc"
// 	query := c.Request.URL.Query()
// 	for key, value := range query {
// 		queryValue := value[len(value)-1]
// 		switch key {
// 		case "per_page":
// 			perPage, _ = strconv.Atoi(queryValue)
// 		case "page":
// 			page, _ = strconv.Atoi(queryValue)
// 		case "sort":
// 			sort = queryValue
// 		}
// 	}

// 	return models.Pagination{
// 		PerPage: perPage,
// 		Page:    page,
// 		Sort:    sort,
// 	}
// }

// var pg = GeneratePaginationFromRequest(c)
// var count int64
// fmt.Println(pg)
// var library = []models.Library{}

// offset := (pg.Page - 1) * pg.PerPage

// queryBuilder := conn.Limit(pg.PerPage).Offset(offset).Order(pg.Sort)
// conn.Model(library).Count(&count)

// queryBuilder.Model(&models.Library{}).Find(&library)

// //Total Page count
// total := int(math.Ceil(float64(count) / float64(pg.PerPage)))
// // sort := "created_at asc"
