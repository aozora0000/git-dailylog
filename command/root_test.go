package command

import "testing"

func TestGetRoot(t *testing.T) {
	t.Run("Test Get Git Root Dir", func(t *testing.T) {
		root, err := getRoot()
		if err != nil {
			t.Error(err.Error())
		} else {
			t.Log(root)
		}
	})
}
