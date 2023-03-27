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

	ConstGetWalletByCustomerID = `
		SELECT 
		    id,
			customer_id, 
		    is_enabled,
			balance,
			create_time,
			created_by,
		    update_time,
		    updated_by
		FROM 
		    wla_wallet 
		WHERE 
		    customer_id = $1
	`

	ConstUpdateWalletStatusByCustomerID = `
		UPDATE 
		    wla_wallet 
		SET 
		    is_enabled = :is_enabled,
		    update_time = :update_time,
		    updated_by = :updated_by
		WHERE 
		    customer_id = :customer_id
	`
)
