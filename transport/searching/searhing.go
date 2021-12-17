package searching

import (
	"context"
"encoding/json"
"fmt"
"net/http"

	"github.com/Petagonest/Check-for-Go/logging"
	"github.com/Petagonest/Check-for-Go/service/search"
	"github.com/julienschmidt/httprouter"
)


func Search(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var search = ps.ByName("search")
	stores, err := search(ctx, search)

	if err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		logging.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	logging.ResponseJSON(w, stores, http.StatusOK)
}
