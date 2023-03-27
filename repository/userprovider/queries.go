package userprovider

const (
	ConstUpsertUserProviderToken = `
		INSERT INTO wla_user_provider (
			customer_id, 
			token, 
			create_time,
			created_by
		) VALUES (
			:customer_id, 
			:token, 
			:create_time, 
			:created_by
		)
        ON CONFLICT (customer_id) DO UPDATE SET
            token = :token,
            update_time = :update_time,
            updated_by = :updated_by
	`

	ConstInsertUserProviderToken = `
		INSERT INTO wla_user_provider (
			customer_id, 
			token, 
			create_time,
			created_by
		) VALUES (
			:customer_id, 
			:token, 
			:create_time, 
			:created_by
		)
	`

	ConstGetUserProviderByToken = `
		SELECT 
		    customer_id, 
		    token, 
		    create_time, 
		    created_by, 
		    update_time, 
		    updated_by 
		FROM 
		    wla_user_provider 
		WHERE token = $1
	`
)
