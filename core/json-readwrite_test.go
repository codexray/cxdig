package core

type jsonTest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type jsonTest2 struct {
	ID   interface{} `json:"id"`
	Name int         `json:"name"`
}

/*
func TestJsonFile(t *testing.T) {
	test := []jsonTest{
		jsonTest{
			ID:   1,
			Name: "test1",
		},
		jsonTest{
			ID:   2,
			Name: "test2",
		},
		jsonTest{
			ID:   3,
			Name: "test3",
		},
	}
	path := "./testingJSONFile.json"
	err := WriteJSONFile(path, test)
	assert.NoError(t, err)
	var test2 []jsonTest
	err = ReadJSONFile(path, &test2)
	assert.NoError(t, err)
	assert.Len(t, test2, 3)
	assert.Equal(t, 1, test2[0].ID)
	assert.Equal(t, "test1", test2[0].Name)
	assert.Equal(t, 2, test2[1].ID)
	assert.Equal(t, "test2", test2[1].Name)
	assert.Equal(t, 3, test2[2].ID)
	assert.Equal(t, "test3", test2[2].Name)

	var test3 []jsonTest2
	assert.Error(t, ReadJSONFile(path, &test3))
	assert.Error(t, ReadJSONFile("./nofile", &test2))
	assert.Error(t, WriteJSONFile(path+"/nofile", test))
	os.Remove(path)
}*/
