package main

func (e *MongoDBMenu) Close() (err error) {
	err = e.Client.Disconnect(e.Ctx)
	return
}
