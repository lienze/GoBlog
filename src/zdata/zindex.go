package zdata

type IndexPageStruct struct {
	PageTitle string
	IndexData []IndexStruct
}

type IndexStruct struct {
	PageTitle   string
	PostPath    string
	PostTitle   string
	PostProfile string
	PostDate    string
}
