package wallet

const (
	ConstCreateNewWallet = `
		INSERT INTO wla_wallet (
			customer_id, 
			balance,
			create_time,
			created_by
		) VALUES (
			:customer_id, 
			:balance, 
			:create_time, 
			:created_by
		)
	`

	ConstUpdateWalletAmount = `
		UPDATE 
		    wla_wallet 
		SET 
		    balance = :balance,
		    update_time = :update_time,
		    updated_by = :updated_by
		WHERE 
		    customer_id = :customer_id
	`

	ConstIncrWalletAmount = `
		UPDATE 
		    wla_wallet 
		SET 
		    balance = balance + :balance,
		    update_time = :update_time,
		    updated_by = :updated_by
		WHERE 
		    customer_id = :customer_id
	`

	ConstGetWalletAmountByCustomerID = `
		SELECT 
		    balance
		FROM 
		    wla_wallet 
		WHERE 
		    customer_id = $1
	`
)
