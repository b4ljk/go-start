package structs

type (
	// Response struct
	Response struct {
		StatusCode int
		Body       ResponseBody
	}

	// ResponseBody struct
	ResponseBody struct {
		Message string      `json:"message"`
		Body    interface{} `json:"body"`
	}

	Sorter map[string]string

	// PaginationInput struct
	PaginationInput struct {
		Limit  int32  `json:"limit"`
		Offset int32  `json:"offset"`
		Sorter Sorter `json:"sorter"`
	}

	//PaginationResponse
	PaginationResponse struct {
		Total int         `json:"total"`
		Items interface{} `json:"items"`
	}

	// CursorInput
	CursorInput struct {
		Limit      int  `form:"limit" json:"limit"`
		PreviousID uint `form:"previous_id" json:"previous_id"`
		HasAll     bool `form:"has_all" json:"has_all"`
	}

	// CursorResponse
	CursorResponse struct {
		HasNext bool        `json:"has_next"`
		Items   interface{} `json:"items"`
	}

	// SuccessResponse
	SuccessResponse struct {
		Success bool `json:"success"`
	}

	// FileInput struct
	FileInput struct {
		ID           int     `json:"id"`
		FileName     string  `json:"file_name"`
		OriginalName string  `json:"original_name"`
		PhysicalPath string  `json:"physical_path"`
		Extention    string  `json:"extention"`
		FileSize     float64 `json:"file_size"`
	}

	ErrorResponse struct {
		StatusCode int    `json:"status_code"`
		ErrorMsg   string `json:"error_msg"`
		Body       string `json:"body"`
	}

	FilterInput struct {
		StartDate *string `json:"start_date"`
		EndDate   *string `json:"end_date"`
		Query     *string `json:"query"`
		IsAll     bool    `json:"is_all"`
	}
)
