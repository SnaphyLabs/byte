package collections

import "github.com/SnaphyLabs/SnaphyByte/models"

type (
	//Interface defining base model list type.
	ModelListProvider interface {
		//Fetch more data from database..
		loadMore() (error)
	}




	// BaseModelList represents a list of items
	BaseModelList struct {
		// Total defines the total number of items in the collection matching the current
		// context. If the storage handler cannot compute this value, -1 is set.
		Total int
		// Offset is the index of the first item of the list in the global collection.
		Offset int
		// Limit is the max number of items requested.
		Limit int
		// ModelProviderList is the list of items contained in the current page given the current
		// context.
		//It doesnot store all the data. Only store current data which has been fetched in the last query..
		ModelProviderList []*models.ModelProvider
	}

)




func (l *BaseModelList) loadMore() error{
	//TODO: fetch data from server..clear the model provider list and add new added data..
	//TODO: Work in progress..

	return nil
}

