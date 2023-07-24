package main



type Item struct {
	title string
	body string
}

var database []Item;


//Read
func GetByName(title string) Item {
	var getItem Item

	for _, val := range database{
		if val.title == title {
			getItem = val
		}
	}

	return getItem
}

//Create
func CreateItem(item Item) Item {
	database = append(database, item)
	return item
}

//Update
func EditItem(title string, edit Item) Item {

	var changed Item

	for idx, val := range database {
		if val.title == edit.title {
			database[idx] = edit
			changed = edit
		}
	}

	return changed;
}

//Delete
func DeleteItem(item Item) Item {
	var deleted Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			deleted =  item
			break
		}
	}

	return deleted
}

func main(){
	//1. functions need to be method
	//2. functions need to be exported: capital first letter on function name
	//3. function need to have two arguments, both are exported types
	//4. the second argument needs to be of a pointer type
	//5. the function return type needs to be Error type
}