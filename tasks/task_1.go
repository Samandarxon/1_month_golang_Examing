package tasks

import (
	"log"
	"playground/newProject/config"
	"playground/newProject/handler"
	"playground/newProject/models"
	"playground/newProject/pkg"
	"playground/newProject/storage"
	"time"

	"github.com/spf13/cast"

	"sort"
)

type tasks struct {
	strg storage.StorageI
	cfg  config.Config
}
type top struct {
	// BranchId string
	Count int
	Name  string
}
type branch struct {
	Summa      int
	BranchName string
}
type eachBranch struct {
	Categories []category
	BranchName string
}

type category struct {
	CategoryName     string
	TransactionCount int
}
type product struct {
	ID       string
	Day      string
	PriceDay int
}
type userEachDay struct {
	User     string
	Products []eachDayProduct
}
type eachDayProduct struct {
	Day     string
	Price   int
	Trcount int
}
type plusMinuse struct {
	BranchName  string
	Transaction transaction
	Summ        summ
}
type transaction struct {
	Pluse  int
	Minuse int
}
type summ struct {
	Pluse  int
	Minuse int
}
type branchPMT struct {
	productId string
	quantity  int
	typed     string
}
type dayProduct struct {
	Day   string
	Count int
}
type includedRemovedProduct struct {
	Name     string
	Included int
	Removed  int
}
type userAmount struct {
	Summa    int
	UserName string
}

var fileNames = models.FileNames{
	BranchFile:                   "data/branches.json",
	UserFile:                     "data/users.json",
	CategoryFile:                 "data/categories.json",
	ProductFile:                  "data/products.json",
	BranchProductFile:            "data/branch_products.json",
	BranchProductTransactionFile: "data/branch_pr_transaction.json",
}

func NewTasksController(strg storage.StorageI, cfg config.Config) *tasks {
	return &tasks{
		strg: strg,
		cfg:  cfg,
	}
}

/* task 12*/
func (t *tasks) TopUserIncludedRemovedProducts() []includedRemovedProduct {

	// fmt.Println("TopIncludedRemovedProducts funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var userAndProductMP = make(map[string][]struct {
		typed    string
		quantity int
	})

	var productMP = make(map[string]includedRemovedProduct)
	var slc = []includedRemovedProduct{}

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		typed := trasn["type"].(string)
		userId := trasn["user_id"].(string)
		quantity := int(trasn["quantity"].(float64))
		// productId := trasn["product_id"].(string)

		userAndProductMP[userId] = append(userAndProductMP[userId], struct {
			typed    string
			quantity int
		}{
			typed:    typed,
			quantity: quantity,
		})
	}
	// fmt.Printf("%+v\n", userAndProductMP)
	for key, val := range userAndProductMP {
		countP := 0
		countM := 0
		for _, v := range val {
			if v.typed == "pluse" {
				countP += v.quantity
			} else {
				countM += v.quantity
			}
		}
		productMP[key] = includedRemovedProduct{
			Name:     h.GetUser(key).Name,
			Included: countP,
			Removed:  countM,
		}
		slc = append(slc, productMP[key])
	}
	// fmt.Printf("%+v\n", slc)

	return slc
}

/* task 11 */
func (t *tasks) TopEachUserEachDaytransaction() []userEachDay {

	// fmt.Println("TopEachBranchNestedEachCategorytransactionSorted funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var mp_transaction = make(map[string]map[string]int)
	var userEachDayMP = make(map[string][]product)
	var slc = []userEachDay{}

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		createdAt := trasn["created_at"].(string)
		productId := trasn["product_id"].(string)
		userId := trasn["user_id"].(string)
		quantity := trasn["quantity"].(float64)
		formatDate1 := "2006-01-02 15:04:05"
		formatDate2 := "2006-01-02"
		date, _ := time.Parse(formatDate1, createdAt)
		createdAt = date.Format(formatDate2)
		// price := int(quantity) * h.GetProduct(productId).Price
		userEachDayMP[userId] = append(userEachDayMP[userId], product{
			ID:       productId,
			Day:      createdAt,
			PriceDay: h.GetProduct(productId).Price * int(quantity),
		})

	}
	// fmt.Println("$$$$$$$$", userEachDayMP, "*************************")
	dayPrice := make(map[string]map[string]int)
	for user_id, v := range userEachDayMP {
		dayPrice[user_id] = make(map[string]int)
		mp_transaction[user_id] = make(map[string]int)
		for _, v2 := range v {
			dayPrice[user_id][v2.Day] += v2.PriceDay
			mp_transaction[user_id][v2.Day]++
		}
	}
	// fmt.Printf("%+v\n", dayPrice)

	// fmt.Printf("$$$$$%+v\n", dar)
	var categoryMap = make(map[string]userEachDay)

	for user_id, val := range dayPrice {

		categories := make([]eachDayProduct, 0)
		for date := range val {
			// fmt.Println(">>>>>%+v\n", v)
			categories = append(categories, eachDayProduct{
				Day:     date,
				Trcount: mp_transaction[user_id][date],
				Price:   dayPrice[user_id][date],
			})
		}
		categoryMap[user_id] = userEachDay{
			User:     h.GetUser(user_id).Name,
			Products: categories,
		}
		slc = append(slc, categoryMap[user_id])
	}

	return slc
}

// task 10
func (t *tasks) TopUserAmountsOfMoney() []userAmount {

	var mp_transaction = make(map[string]int)

	var slc = []userAmount{}

	// fmt.Println("TopMoneytopansactionSorted funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		// typed := trasn["type"].(string)
		userId := trasn["user_id"].(string)
		productId := trasn["product_id"].(string)
		quantity := trasn["quantity"].(float64)
		price := int(quantity) * h.GetProduct(productId).Price
		// if typed == "pluse" {
		mp_transaction[userId] += price
		// }

	}
	// fmt.Println(mp_transaction)
	for key, val := range mp_transaction {
		slc = append(slc, userAmount{
			Summa:    val,
			UserName: h.GetUser(key).Name,
		})
	}
	sort.Slice(slc, func(a, b int) bool {
		return slc[a].Summa > slc[b].Summa
	})

	return slc

}

/* task 9*/
func (t *tasks) TopBranchAmountsOfMoney() []branch {

	var mp_transaction = make(map[string]int)

	var slc = []branch{}

	// fmt.Println("TopMoneytopansactionSorted funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		// typed := trasn["type"].(string)
		branchId := trasn["branch_id"].(string)
		productId := trasn["product_id"].(string)
		quantity := trasn["quantity"].(float64)
		price := int(quantity) * h.GetProduct(productId).Price
		// fmt.Println(branchId, h.GetBranch(branchId).Name, productId, h.GetProduct(productId).Price, h.GetProduct(productId).Name, "price", price, "quantity", quantity)
		// if typed == "pluse" {
		mp_transaction[branchId] += price
		// }

	}
	// fmt.Println(mp_transaction)
	for key, val := range mp_transaction {
		slc = append(slc, branch{
			Summa:      val,
			BranchName: h.GetBranch(key).Name,
		})
	}
	sort.Slice(slc, func(a, b int) bool {
		return slc[a].Summa > slc[b].Summa
	})

	return slc
}

/* task 8*/
func (t *tasks) TopIncludedRemovedProducts() []includedRemovedProduct {

	// fmt.Println("TopIncludedRemovedProducts funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var branchAndProductMP = make(map[string][]struct {
		typed    string
		quantity int
	})

	var productMP = make(map[string]includedRemovedProduct)
	var slc = []includedRemovedProduct{}

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		typed := trasn["type"].(string)
		quantity := int(trasn["quantity"].(float64))
		productId := trasn["product_id"].(string)

		branchAndProductMP[productId] = append(branchAndProductMP[productId], struct {
			typed    string
			quantity int
		}{
			typed:    typed,
			quantity: quantity,
		})
	}
	// fmt.Printf("%+v\n", branchAndProductMP)
	for key, val := range branchAndProductMP {
		countP := 0
		countM := 0
		for _, v := range val {
			if v.typed == "pluse" {
				countP += v.quantity
			} else {
				countM += v.quantity
			}
		}
		productMP[key] = includedRemovedProduct{
			Name:     h.GetProduct(key).Name,
			Included: countP,
			Removed:  countM,
		}
		slc = append(slc, productMP[key])
	}
	// fmt.Printf("%+v\n", slc)

	return slc
}

/* task 7*/
func (t *tasks) TopEachDayProducts() []dayProduct {

	// fmt.Println("TopEachDayProducts funcsiya")
	// h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var branchAndProductMP = make(map[string][]struct {
		typed    string
		quantity int
	})

	var slc = []dayProduct{}

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		typed := trasn["type"].(string)
		quantity := int(trasn["quantity"].(float64))
		createdAt := trasn["created_at"].(string)
		// productId := trasn["product_id"].(string)
		formatDate1 := "2006-01-02 15:04:05"
		formatDate2 := "2006-01-02"
		date, _ := time.Parse(formatDate1, createdAt)
		createdAt = date.Format(formatDate2)
		if typed == "pluse" {

			branchAndProductMP[createdAt] = append(branchAndProductMP[createdAt], struct {
				typed    string
				quantity int
			}{
				typed:    typed,
				quantity: quantity,
			})
		}
	}
	var productMP = make(map[string]dayProduct)
	for key, val := range branchAndProductMP {
		count := 0
		for _, v := range val {
			count += v.quantity
		}
		productMP[key] = dayProduct{
			Count: count,
			Day:   key,
		}
		slc = append(slc, productMP[key])
	}
	// fmt.Printf("%+v\n", slc)

	return slc
}

/* task 6 */
func (t *tasks) TopBranchPluseMinusetransactionSorted() []plusMinuse {

	// fmt.Println("TopBranchPluseMinusesactionSorted funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// var mp_transaction = make(map[string]int)
	var branchAndProductMP = make(map[string][]branchPMT)

	var slc = []plusMinuse{}

	var categoryMap = make(map[string]plusMinuse)

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		branchId := trasn["branch_id"].(string)
		typed := trasn["type"].(string)
		quantity := int(trasn["quantity"].(float64))
		productId := trasn["product_id"].(string)
		branchAndProductMP[branchId] = append(branchAndProductMP[branchId], branchPMT{
			productId: productId,
			typed:     typed,
			quantity:  quantity,
		})
	}
	for key, val := range branchAndProductMP {
		// fmt.Println(">>>>>>>>", key, val)
		tr := transaction{}
		sm := summ{}
		for _, val2 := range val {
			// fmt.Println("quantity ", val2)
			if val2.typed == "pluse" {
				tr.Pluse += val2.quantity
				sm.Pluse += h.GetProduct(val2.productId).Price * val2.quantity
			} else {
				tr.Minuse += val2.quantity
				sm.Minuse += h.GetProduct(val2.productId).Price * val2.quantity
			}
		}
		categoryMap[key] = plusMinuse{
			BranchName:  h.GetBranch(key).Name,
			Transaction: tr,
			Summ:        sm,
		}

		slc = append(slc, categoryMap[key])

	}

	return slc
}

/* task 5 */
func (t *tasks) TopEachBranchNestedEachCategorytransaction() []eachBranch {

	// fmt.Println("TopEachBranchNestedEachCategorytransactionSorted funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var mp_transaction = make(map[string]map[string]int)
	var branchAndProductMP = make(map[string][]product)
	var slc = []eachBranch{}

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		branchId := trasn["branch_id"].(string)
		productId := trasn["product_id"].(string)
		branchAndProductMP[branchId] = append(branchAndProductMP[branchId], product{
			ID: productId,
		})

	}
	for key, val := range branchAndProductMP {
		mp_transaction[key] = make(map[string]int)
		for _, val2 := range val {
			categoryId := h.GetProduct(val2.ID).CategoryId
			mp_transaction[key][categoryId]++
		}

	}
	var categoryMap = make(map[string]eachBranch)

	for key, val := range mp_transaction {

		categories := make([]category, 0)
		for key1, val1 := range val {

			categories = append(categories, category{
				CategoryName:     h.GetCategory(key1).Name,
				TransactionCount: val1,
			})
		}
		categoryMap[key] = eachBranch{
			BranchName: h.GetBranch(key).Name,
			Categories: categories,
		}
		slc = append(slc, categoryMap[key])
	}

	return slc
}

/* task 4 */
func (t *tasks) TopCategorytransactionSorted() []top {

	// fmt.Println("TopCategorysactionSorted funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var mp_transaction = make(map[string]int)

	var slc = []top{}

	var categoryMap = make(map[string]top)

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		productId := trasn["product_id"].(string)
		categoryId := h.GetProduct(productId).CategoryId
		mp_transaction[categoryId]++
	}

	for key, val := range mp_transaction {

		categoryMap[key] = top{
			Count: val,
			Name:  h.GetCategory(key).Name,
		}
		slc = append(slc, categoryMap[key])

	}

	sort.Slice(slc, func(a, b int) bool {
		return slc[a].Count > slc[b].Count
	})

	return slc
}

// task 3
func (t *tasks) TopProductransactionSorted() []top {

	// fmt.Println("TopProductransactionSorted funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var mp_transaction = make(map[string]int)

	var slc = []top{}

	var productMap = make(map[string]top)

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		productId := trasn["product_id"].(string)
		mp_transaction[productId]++

	}

	for key, val := range mp_transaction {

		productMap[key] = top{
			Count: val,
			Name:  h.GetProduct(key).Name,
		}
		slc = append(slc, productMap[key])

	}

	sort.Slice(slc, func(a, b int) bool {
		return slc[a].Count > slc[b].Count
	})

	return slc
}

// task 2
func (t *tasks) TopMoneytopansactionSorted() []branch {
	var mp_transaction = make(map[string]int)

	var slc = []branch{}

	// fmt.Println("TopMoneytopansactionSorted funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		branchId := trasn["branch_id"].(string)
		productId := trasn["product_id"].(string)
		quantity := trasn["quantity"].(float64)
		price := int(quantity) * h.GetProduct(productId).Price
		// fmt.Println(branchId, h.GetBranch(branchId).Name, productId, h.GetProduct(productId).Price, h.GetProduct(productId).Name, "price", price, "quantity", quantity)
		mp_transaction[branchId] += price

	}
	// fmt.Println(mp_transaction)
	for key, val := range mp_transaction {
		slc = append(slc, branch{
			Summa:      val,
			BranchName: h.GetBranch(key).Name,
		})
	}
	sort.Slice(slc, func(a, b int) bool {
		return slc[a].Summa > slc[b].Summa
	})

	return slc
}

// task 1
func (t *tasks) TopBranchTransactionSorted() []top {

	// fmt.Println("ToptopansactionSorted funcsiya")
	h := handler.NewHandler(t.strg, t.cfg)

	data_trasaction, err := pkg.Read(fileNames.BranchProductTransactionFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var mp_transaction = make(map[string]int)

	var slc = []top{}

	var branchMap = top{}

	for _, val := range data_trasaction {
		var trasn = cast.ToStringMap(val)
		branchId := trasn["branch_id"].(string)
		mp_transaction[branchId]++

	}

	for key, val := range mp_transaction {
		branchMap = top{
			// BranchId: key,
			Count: val,
			Name:  h.GetBranch(key).Name,
		}
		slc = append(slc, branchMap)

	}

	sort.Slice(slc, func(a, b int) bool {
		return slc[a].Count > slc[b].Count
	})

	return slc
}
