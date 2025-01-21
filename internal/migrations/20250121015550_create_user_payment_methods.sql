-- +goose Up
CREATE TABLE user_payment_methods (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    uuid CHAR(36) NOT NULL,

    -- Loại phương thức thanh toán
    payment_type ENUM('bank', 'card', 'e_wallet') NOT NULL,
    provider VARCHAR(50) NOT NULL,
    
    -- Thông tin chung
    is_default BOOLEAN DEFAULT FALSE,
    status ENUM('active', 'inactive') DEFAULT 'active',
    
    -- Thông tin chi tiết (tùy loại)
    card_info JSON,
    bank_info JSON,
    wallet_info JSON,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    -- Các Constraints và Indexes
    CONSTRAINT fk_user_payment_methods_user_id 
        FOREIGN KEY (user_id) 
        REFERENCES user_basics(id) 
        ON DELETE CASCADE,
    
    UNIQUE KEY unique_uuid (uuid),
        
    INDEX idx_user_id (user_id),
    INDEX idx_user_default (user_id, is_default),
    INDEX idx_payment_type (payment_type),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
DROP TABLE IF EXISTS user_payment_methods;