package repositories

import (
	"fmt"
	"white-goods-multifinace/dto"
	"white-goods-multifinace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchaseRepository interface {
	CreatePurchase(purchase *models.Purchase) error
	FindPurchaseByID(purchaseID uuid.UUID) (*dto.PurchaseByIDResponse, error)
	FindAllPurchase(userID uuid.UUID) (*[]dto.GetAllUserPurchase, error)
}

type purchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) PurchaseRepository {
	return &purchaseRepository{db: db}
}

func (r *purchaseRepository) CreatePurchase(purchase *models.Purchase) error {
	if err := r.db.Create(purchase).Error; err != nil {
		return err
	}

	return nil
}

func (r *purchaseRepository) FindAllPurchase(userID uuid.UUID) (*[]dto.GetAllUserPurchase, error) {
	var purchases []dto.GetAllUserPurchase

	// Query untuk mendapatkan data Purchase dan relasi terkait
	query := `
		SELECT 
			p.id AS purchase_id, 
			p.monthly_payment, 
			p.is_completed, 
			it.id AS item_tenor_id, 
			it.interest, 
			i.id AS item_id, 
			i.name AS item_name, 
			i.normal_price, 
			i.admin_fee, 
			o.id AS otr_id, 
			o.otr AS otr_name, 
			t.id AS tenor_id, 
			t.duration AS tenor_duration
		FROM purchases p
		JOIN user_limits ul ON ul.id = p.user_limit_id
		JOIN item_tenors it ON p.item_tenor_id = it.id
		JOIN items i ON it.item_id = i.id
		JOIN tenors t ON it.tenor_id = t.id
		LEFT JOIN otrs o ON i.otr_id = o.id
		WHERE ul.user_id = ?`

	rows, err := r.db.Raw(query, userID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var purchase dto.GetAllUserPurchase
		var itemTenor dto.GetAllPurchaseItemTenor
		var item dto.ItemResponse
		var tenor dto.TenorResponse
		var otr dto.OTRResponse

		err := rows.Scan(
			&purchase.PurchaseID, &purchase.MonthlyPayment, &purchase.IsCompleted,
			&itemTenor.ItemTenorID, &itemTenor.Interest,
			&item.ItemID, &item.Name, &item.NormalPrice, &item.AdminFee,
			&otr.OTRID, &otr.Name,
			&tenor.TenorID, &tenor.Duration,
		)
		if err != nil {
			return nil, err
		}

		itemTenor.Item = item
		itemTenor.Tenor = tenor
		purchase.Purchases = itemTenor

		purchases = append(purchases, purchase)
	}

	if len(purchases) == 0 {
		return nil, fmt.Errorf("no purchases found for user")
	}

	return &purchases, nil
}

func (r *purchaseRepository) FindPurchaseByID(purchaseID uuid.UUID) (*dto.PurchaseByIDResponse, error) {
	var purchase models.Purchase
	var itemTenor models.ItemTenor
	var item models.Item
	var tenor models.Tenor
	var otr models.OTR
	var transaction []models.Transaction
	var transactionResponses []dto.TransactionResponse

	query := `
		SELECT 
			p.id AS purchase_id, 
			p.monthly_payment, 
			p.is_completed, 
			it.id AS item_tenor_id, 
			it.interest, 
			i.id AS item_id, 
			i.name AS item_name, 
			i.normal_price, 
			i.admin_fee, 
			o.id AS otr_id, 
			o.otr AS otr_name, 
			t.id AS tenor_id, 
			t.duration AS tenor_duration
		FROM purchases p
		JOIN item_tenors it ON p.item_tenor_id = it.id
		JOIN items i ON it.item_id = i.id
		JOIN tenors t ON it.tenor_id = t.id
		LEFT JOIN otrs o ON i.otr_id = o.id
		WHERE p.id = ?`

	rows, err := r.db.Raw(query, purchaseID).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&purchase.ID, &purchase.MonthlyPayment, &purchase.IsCompleted,
			&itemTenor.ID, &itemTenor.Interest,
			&item.ID, &item.Name, &item.NormalPrice, &item.AdminFee,
			&otr.ID, &otr.OTR,
			&tenor.ID, &tenor.Duration,
		)
		if err != nil {
			return nil, err
		}
	}

	if err := r.db.Find(&transaction, "purchase_id = ?", purchaseID).Error; err != nil {
		return nil, err
	}

	for _, tx := range transaction {
		transactionResponses = append(transactionResponses, dto.TransactionResponse{
			TransactionID: tx.ID.String(),
			TotalAmount:   tx.TotalAmount,
			PaymentDate:   tx.PaymentDate,
			InvoiceNumber: tx.InvoiceNumber,
		})
	}

	response := dto.PurchaseByIDResponse{
		PurchaseID:     purchase.ID,
		MonthlyPayment: purchase.MonthlyPayment,
		IsCompleted:    purchase.IsCompleted,
		ItemTenor: dto.ItemTenorResponse{
			ItemTenorID: itemTenor.ID,
			Interest:    itemTenor.Interest,
			Item: dto.ItemResponse{
				ItemID:      item.ID,
				Name:        item.Name,
				NormalPrice: item.NormalPrice,
				AdminFee:    item.AdminFee,
				OTR: dto.OTRResponse{
					OTRID: otr.ID,
					Name:  otr.OTR,
				},
			},
			Tenor: dto.TenorResponse{
				TenorID:  tenor.ID,
				Duration: int(tenor.Duration),
			},
			Transactions: transactionResponses,
		},
	}

	return &response, nil
}
