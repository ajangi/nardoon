package utils

// GetErrorByKey : A method to get all locales errors for response messages by key
func GetErrorByKey(key string) MessageItem {
	messages := map[string]MessageItem{
		NotFoundErrorMessageKey: MessageItem{Fa: "یافت نشد!", En: "Ooops! Not found..."},
	}
	suitableMessage := messages[key]
	if (suitableMessage == MessageItem{}) {
		return MessageItem{Fa: key, En: key}
	}
	return suitableMessage
}
