package storeformat

//Store is the format for the storage in memstore
type Store struct {
	KeyValue map[string]interface{}
	List     []interface{}
}
