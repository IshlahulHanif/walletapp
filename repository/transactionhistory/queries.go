package transactionhistory

const (
	ConstInsertNewTransaction = `
		INSERT INTO wla_wallet_transaction_history (
			status, 
			wallet_id, 
			transacted_at, 
			type, 
			amount, 
			reference_id
		) VALUES (
			:status, 
			:wallet_id, 
			:transacted_at, 
			:type, 
			:amount, 
			:reference_id
		)
	`

	ConstGetAllTransactionByWalletID = `
		SELECT 
		    id, 
		    status, 
		    wallet_id, 
		    transacted_at, 
		    type, 
		    amount, 
		    reference_id 
		FROM 
		    wla_wallet_transaction_history 
		WHERE 
		    wallet_id = $1
		ORDER BY 
		    transacted_at DESC
	`
)
