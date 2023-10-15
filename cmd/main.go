package main

import (
	"fmt"
	"playground/newProject/config"
	"playground/newProject/models"
	"playground/newProject/storage"
	"playground/newProject/storage/memory"
	"playground/newProject/tasks"
	"strconv"
	"strings"
)

func main() {

	cfg := config.Load()
	strg := memory.NewStorage(models.FileNames{
		BranchFile:                   "data/branches.json",
		UserFile:                     "data/users.json",
		CategoryFile:                 "data/categories.json",
		ProductFile:                  "data/products.json",
		BranchProductFile:            "data/branch_products.json",
		BranchProductTransactionFile: "data/branch_pr_transaction.json",
	})
	// h := handler.NewHandler(strg, *cfg)
	// h.CreateCategory("FastFood")
	// h.GetAllBranchProductTransaction(1, 10, "", "")
	// h.CreateBranchProductTransaction("f7497e5f-ce02-452b-8ea4-d5aab1c69a1c", "b9e54f3a-7e11-4b0a-94c3-b7e27e601044", "904d377a-75e4-4dd0-a80d-9449e7b91677", "pluse", 2)
	// h.CreateBranch("Axorazmiy", "Chilonzor")

	// Tasks

	// task_12(strg, *cfg)
	task_11(strg, *cfg)
	// task_10(strg, *cfg)
	// task_9(strg, *cfg)
	// task_8(strg, *cfg)
	// task_7(strg, *cfg)
	// task_6(strg, *cfg)
	// task_5(strg, *cfg)
	// task_4(strg, *cfg)
	// task_3(strg, *cfg)
	// task_2(strg, *cfg)
	// task_1(strg, *cfg)

	// for index, val := range tasks.TopEachBranchNestedEachCategorytransactionSorted() {
	// 	fmt.Printf("%d %s %v\n", index+1, val.BranchName, val.Categories)
	// }
	// fmt.Println(tasks.TopEachBranchNestedEachCategorytransactionSorted())
	// Category
	// h.CreateCategory("Sabzavod")
	// h.UpdateCategory("eee3d471-78c1-4aaf-8c62-4df8acbb5c1f", "Torix")
	// h.DeleteCategory("eee3d471-78c1-4aaf-8c62-4df8acbb5c1f")
	// h.GetAllCategory(1, 10)
	// h.GetCategory("ff4cb227-92d7-4a7f-83fb-4d7e3e52d4aa")
	// Product
	// h.CreateProduct("52675e16-f992-41ae-9f8d-ec2fd8c77e44", "Sok", 10500)
	// h.UpdateProduct("e4b2c0f4-b0b5-43ec-a48b-f66513e73d66", "b8301a79-a369-40b5-b518-3d68d2b43c38", "Kasha", 100_000)
	// h.GetAllProduct(1, 10)
	// h.GetProduct("e4b2c0f4-b0b5-43ec-a48b-f66513e73d66")
	// h.DeleteProduct("e4b2c0f4-b0b5-43ec-a48b-f66513e73d66")
}

// task 12
func task_12(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	maxLength := 30
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*2)))
	fmt.Printf("\t Name\t\t\t  Included\tRemoved\n")
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*2)))
	for index, val := range tasks.TopUserIncludedRemovedProducts() {
		str := fmt.Sprint(strings.Repeat(" ", maxLength-len(val.Name)))

		fmt.Printf("  %d\t%s%s%d%s%d\n", index+1,
			val.Name,
			str,
			val.Included,
			fmt.Sprint(strings.Repeat(" ", maxLength/2-len(strconv.Itoa(val.Included)))),
			val.Removed,
		)
	}
}

// task 11
func task_11(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	maxLength := 30
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3)))
	fmt.Printf("\tName%sDate%sTrcount%sPrice\n",
		fmt.Sprint(strings.Repeat(" ", maxLength-13-len("Name"))),
		fmt.Sprint(strings.Repeat(" ", maxLength-13-len("Date"))),
		fmt.Sprint(strings.Repeat(" ", maxLength-13-len("Price"))),
	)
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3)))
	for index, val := range tasks.TopEachUserEachDaytransaction() {
		// fmt.Printf("%d  \t %+v\n", index+1, val)

		maxLength = 20
		str := fmt.Sprint(strings.Repeat(" ", maxLength-len(val.User)))

		fmt.Printf("  %d\t%s%s", index+1, val.User, str)
		i := 1
		for _, v := range val.Products {
			if len(val.Products) > 1 {
				cn := v.Day
				if i > 1 {

					str = fmt.Sprint(strings.Repeat(" ", maxLength+len(cn)-3))
					fmt.Printf(" %s%s%s%d%s%d\n",
						str,
						v.Day,
						fmt.Sprint(strings.Repeat(" ", maxLength-len(cn)-len(strconv.Itoa(v.Trcount)))),
						v.Trcount,
						// str,
						fmt.Sprint(strings.Repeat(" ", maxLength-len(cn)-len(strconv.Itoa(v.Trcount)))),
						v.Price)
				} else {
					str = fmt.Sprint(strings.Repeat(" ", maxLength-len(cn)-len(strconv.Itoa(v.Trcount))))
					fmt.Printf("%s%s%d%s%d\n", v.Day, str, v.Trcount, str, v.Price)
				}
				i++
				continue
			} else {
				str = fmt.Sprint(strings.Repeat(" ", maxLength-len(v.Day)-len(strconv.Itoa(v.Trcount))))
				fmt.Printf("%s%s%d%s%d\n", v.Day, str, v.Trcount, str, v.Price)
			}
		}
	}
	// fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3)))
	// fmt.Println()

}

// task 10
func task_10(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	maxLength := 20
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*2+10)))
	fmt.Printf("\tName%sAmount\n", fmt.Sprint(strings.Repeat(" ", maxLength)))
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*2+10)))

	for index, val := range tasks.TopUserAmountsOfMoney() {
		fmt.Printf("%d\t%s%s\t%d сум\n", index+1,
			val.UserName,
			fmt.Sprint(strings.Repeat(" ", maxLength-len(val.UserName))),
			val.Summa,
		)
	}
}

// task 9
func task_9(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	maxLength := 20
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*2+10)))
	fmt.Printf("\tName%sAmount\n", fmt.Sprint(strings.Repeat(" ", maxLength)))
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*2+10)))

	for index, val := range tasks.TopBranchAmountsOfMoney() {
		fmt.Printf("%d\t%s%s\t%d сум\n", index+1,
			val.BranchName,
			fmt.Sprint(strings.Repeat(" ", maxLength-len(val.BranchName))),
			val.Summa,
		)
	}
}

// task 8
func task_8(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	maxLength := 30
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*2)))
	fmt.Printf("\t Name\t\t\t  Included\tRemoved\n")
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*2)))
	for index, val := range tasks.TopIncludedRemovedProducts() {
		str := fmt.Sprint(strings.Repeat(" ", maxLength-len(val.Name)))

		fmt.Printf("  %d\t%s%s%d%s%d\n", index+1,
			val.Name,
			str,
			val.Included,
			fmt.Sprint(strings.Repeat(" ", maxLength/2-len(strconv.Itoa(val.Included)))),
			val.Removed,
		)
	}
}

// task 7
func task_7(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	maxLength := 10
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3+5)))
	fmt.Printf("\t Day\t\tCount\n")
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3+5)))

	for index, val := range tasks.TopEachDayProducts() {

		fmt.Printf("  %d   %s\t  %d\n", index+1, val.Day, val.Count)
	}
}

// task 6
func task_6(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	maxLength := 40
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3)))
	fmt.Printf("\tBranchName%sTransactions%sSumma\n",
		fmt.Sprint(strings.Repeat(" ", maxLength-len("Transactions")-4)),
		fmt.Sprint(strings.Repeat(" ", maxLength-len("Summa"))),
	)

	// fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3)))

	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3)))
	fmt.Printf("%sPluse%sMinuse%sPluse%sMinuse\n\n",
		fmt.Sprint(strings.Repeat(" ", maxLength)),
		fmt.Sprint(strings.Repeat(" ", len("Pluse"))),
		fmt.Sprint(strings.Repeat(" ", maxLength-len("Pluse")-10)),
		fmt.Sprint(strings.Repeat("  ", len("Minuse"))),
	)
	// fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3)))

	for index, val := range tasks.TopBranchPluseMinusetransactionSorted() {
		str := fmt.Sprint(strings.Repeat(" ", maxLength-len(val.BranchName)-5))
		fmt.Printf("%d\t%s%s%d%s%d%s%d%s%d\n", index+1,
			val.BranchName,
			str,
			val.Transaction.Pluse,
			fmt.Sprint(strings.Repeat(" ", maxLength/4-len(strconv.Itoa(val.Transaction.Pluse)))),
			val.Transaction.Minuse,
			fmt.Sprint(strings.Repeat(" ", maxLength-12-len(strconv.Itoa(val.Transaction.Minuse)))),
			val.Summ.Pluse,
			fmt.Sprint(strings.Repeat(" ", maxLength/3+5-len(strconv.Itoa(val.Summ.Pluse)))),

			val.Summ.Minuse,
		)
	}
}

// task 5
func task_5(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	maxLength := 30
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3)))
	fmt.Printf("\tBranchName%sCategoryName%sTransactionCount\n", fmt.Sprint(strings.Repeat(" ", maxLength-len("CategoryName"))), fmt.Sprint(strings.Repeat(" ", 5+maxLength-len("TransactionCount"))))
	fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3)))
	for index, val := range tasks.TopEachBranchNestedEachCategorytransaction() {

		str := fmt.Sprint(strings.Repeat(" ", maxLength-len(val.BranchName)))

		fmt.Printf("  %d\t%s%s", index+1, val.BranchName, str)
		i := 1
		for _, v := range val.Categories {
			// maxLength = 50
			if len(val.Categories) > 1 {
				cn := v.CategoryName
				if i > 1 {

					str = fmt.Sprint(strings.Repeat(" ", maxLength+len(cn)))
					fmt.Printf(" %s%s%s%d\n", str, v.CategoryName, fmt.Sprint(strings.Repeat(" ", maxLength-len(cn)-len(strconv.Itoa(v.TransactionCount)))), v.TransactionCount)
				} else {
					str = fmt.Sprint(strings.Repeat(" ", maxLength-len(cn)-len(strconv.Itoa(v.TransactionCount))))
					fmt.Printf("%s%s%d\n", v.CategoryName, str, v.TransactionCount)
				}
				i++
				continue
			} else {
				str = fmt.Sprint(strings.Repeat(" ", maxLength-len(v.CategoryName)-len(strconv.Itoa(v.TransactionCount))))
				fmt.Printf("%s%s%d\n", v.CategoryName, str, v.TransactionCount)
			}
		}
		// fmt.Println(fmt.Sprint(strings.Repeat("-", maxLength*3)))
		// fmt.Println()
	}
}

// task 4
func task_4(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	fmt.Println("\tCategoryName\t\tTransactionCount")

	maxLength := 30
	for index, val := range tasks.TopCategorytransactionSorted() {
		str := fmt.Sprint(strings.Repeat(" ", maxLength-len(val.Name)))
		fmt.Printf("%d\t%s%s%d\n", index+1, val.Name, str, val.Count)
	}
}

// task 3
func task_3(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	fmt.Println("\tProductName\t\tTransactionCount")

	maxLength := 30
	for index, val := range tasks.TopProductransactionSorted() {
		str := fmt.Sprint(strings.Repeat(" ", maxLength-len(val.Name)))
		fmt.Printf("%d\t%s%s%d\n", index+1, val.Name, str, val.Count)
	}
}

// task 2
func task_2(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	maxLength := 15
	fmt.Printf("\tBranchName%sSumma\n", fmt.Sprint(strings.Repeat(" ", maxLength-len("Summa"))))
	maxLength = 20

	for index, val := range tasks.TopMoneytopansactionSorted() {
		str := fmt.Sprint(strings.Repeat(" ", maxLength-len(val.BranchName)))
		fmt.Printf("%d\t%s%s%d сум\n", index+1, val.BranchName, str, val.Summa)
	}
}

// task 1
func task_1(strg storage.StorageI, cfg config.Config) {
	tasks := tasks.NewTasksController(strg, cfg)
	fmt.Println("\tBranchName\t\tTransactionCount")

	maxLength := 30
	for index, val := range tasks.TopBranchTransactionSorted() {
		str := fmt.Sprint(strings.Repeat(" ", maxLength-len(val.Name)))
		fmt.Printf("%d\t%s%s%d\n", index+1, val.Name, str, val.Count)
	}
}
