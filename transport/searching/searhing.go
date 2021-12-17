package searching

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Petagonest/Check-for-Go/logging"
	search "github.com/Petagonest/Check-for-Go/service/searching"
	"github.com/julienschmidt/httprouter"
)

func Search(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var SearchingAll = ps.ByName("search")
	search1, search2, search3, err := search.SearchingAll(ctx, SearchingAll, SearchingAll, SearchingAll)

	if err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		logging.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	logging.ResponseJSONSEARCH(w, search1, search2, search3, http.StatusOK)
}
