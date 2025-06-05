
from typing import List

# Given a list of accounts where each element accounts[i] is a list of 
# strings, where the first element accounts[i][0] is a name,
#  and the rest of the elements are emails representing emails of the account.

# Now, we would like to merge these accounts.
#  Two accounts definitely belong to the same person
#  if there is some common email to both accounts.
#  Note that even if two accounts have the same name,
#  they may belong to different people as people could have the same name.
#  A person can have any number of accounts initially,
#  but all of their accounts definitely have the same name.

# After merging the accounts, return the accounts in the following format:
#  the first element of each account is the name, and
#  the rest of the elements are emails in sorted order.
#  The accounts themselves can be returned in any order.

# Example 1:

# Input: accounts = [
#     ["neet","neet@gmail.com","neet_dsa@gmail.com"],
#     ["alice","alice@gmail.com"],
#     ["neet","bob@gmail.com","neet@gmail.com"],
#     ["neet","neetcode@gmail.com"]
# ]

# Output: [["neet","bob@gmail.com","neet@gmail.com","neet_dsa@gmail.com"],["alice","alice@gmail.com"],["neet","neetcode@gmail.com"]]
# Example 2:

# Input: accounts = [
#     ["James","james@mail.com"],
#     ["James","james@mail.co"]
# ]

# Output: [["James","james@mail.com"],["James","james@mail.co"]]

class Solution:
    def accountsMerge(self, accounts: List[List[str]]) -> List[List[str]]:
        # Step 1: Build connections between accounts
        graph = [[] for _ in range(len(accounts))]

        # 🚨 MISSING: Find which accounts share emails
        email_to_accounts = {}
        for i, account in enumerate(accounts):
            for email in account[1:]:  # Skip name
                if email in email_to_accounts:
                    # Connect current account to all previous accounts with this email
                    for prev_account in email_to_accounts[email]:
                        graph[i].append(prev_account)
                        graph[prev_account].append(i)
                else:
                    email_to_accounts[email] = []
                email_to_accounts[email].append(i)
        
        visited = [False] * len(accounts)
        result = []
        
        for i in range(len(accounts)):
            if not visited[i]:
                # DFS to find all accounts in this component
                component_emails = set()
                name = accounts[i][0]
                
                def dfs(account_idx):
                    visited[account_idx] = True
                    # Add emails from this account
                    for email in accounts[account_idx][1:]:
                        component_emails.add(email)
                    # Visit connected accounts
                    for neighbor in graph[account_idx]:
                        if not visited[neighbor]:
                            dfs(neighbor)
                
                dfs(i)
                # Build result for this component
                result.append([name] + sorted(component_emails))
        
        return result


accounts = [["neet","neet@gmail.com","neet_dsa@gmail.com"],
    ["alice","alice@gmail.com"],
    ["neet","bob@gmail.com","neet@gmail.com"],
    ["neet","neetcode@gmail.com"]]