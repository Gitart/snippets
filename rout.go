package main

func Srouter(Y string) string {

	// Таблицы приемники
	T := []string{"Directory", "Other", "Docarchive", "Docbackup"}

	S := map[string]string{
		"Persons":       T[0],
		"Cashreasons":   T[0],
		"Structures":    T[0],
		"Nets":          T[2],
		"Cls":           T[3],
		"FiscalRegistr": T[0],
		"DrugSupplier":  T[0],
		"Aplantempl":    T[0],
	}

	R := ""

	// Добавить процедуру проверки таблицы в базе
	// Если нет - создать

	// Нет данных
	if Y == "" {
		return R
	}

	// Поиск
	for K, P := range S {
		if K == Y {
			R = P
			break
		} else {
			R = Y
		}
	}
