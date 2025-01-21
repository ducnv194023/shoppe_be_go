-- +goose Up
CREATE TABLE user_addresses (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    uuid CHAR(36) NOT NULL,


    receiver_name VARCHAR(255) NOT NULL,
    receiver_phone VARCHAR(15) NOT NULL,
    
    -- Địa chỉ chi tiết
    province_code VARCHAR(20) NOT NULL,
    district_code VARCHAR(20) NOT NULL,
    ward_code VARCHAR(20) NOT NULL,
    street_address TEXT NOT NULL,
    
    -- Thông tin bổ sung
    label VARCHAR(50),           
    is_default BOOLEAN DEFAULT FALSE,
    notes TEXT,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT fk_user_addresses_user_id 
        FOREIGN KEY (user_id) 
        REFERENCES user_basics(id) 
        ON DELETE CASCADE,
    
    UNIQUE KEY unique_uuid (uuid),
    INDEX idx_user_id (user_id),          
    INDEX idx_user_default (user_id, is_default)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
DROP TABLE IF EXISTS user_addresses;