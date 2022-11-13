package model

type PutCSV struct {
	Date string  `csv:"date"`
	Name string  `csv:"name"`
	ID   string  `csv:"id"`
	Age  uint    `csv:"age"`
	Num  float32 `csv:"number"`
}

type InvokeInput struct {
	Key    string
	Method string
}
