package unionfind

// MergeAccounts groups accounts that share common emails
func MergeAccounts(accountList [][]string) [][]string {
	emailToAccountIndex := make(map[string]int)
	accountUnionFind := NewUnionFind(len(accountList))

	// Build union-find by linking accounts that share emails
	for accountIndex, accountEntry := range accountList {
		for emailIndex := 1; emailIndex < len(accountEntry); emailIndex++ {
			emailAddress := accountEntry[emailIndex]
			if existingAccount, emailFound := emailToAccountIndex[emailAddress]; emailFound {
				accountUnionFind.Union(accountIndex, existingAccount)
			} else {
				emailToAccountIndex[emailAddress] = accountIndex
			}
		}
	}

	// Group emails by their root account
	rootToEmails := make(map[int]map[string]bool)
	for emailAddress, accountIndex := range emailToAccountIndex {
		rootAccount := accountUnionFind.Find(accountIndex)
		if rootToEmails[rootAccount] == nil {
			rootToEmails[rootAccount] = make(map[string]bool)
		}
		rootToEmails[rootAccount][emailAddress] = true
	}

	// Build result with account name and sorted emails
	mergedAccounts := [][]string{}
	for rootAccount, emailSet := range rootToEmails {
		accountName := accountList[rootAccount][0]
		emailList := []string{accountName}
		for emailAddress := range emailSet {
			emailList = append(emailList, emailAddress)
		}
		mergedAccounts = append(mergedAccounts, emailList)
	}
	return mergedAccounts
}
