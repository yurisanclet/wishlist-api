package repositories

import (
	"backend/config"
	"backend/models"
	"errors"
)

type WishListRepository interface {
	CreateWishList(wishlist *models.Wishlist) error
	GetWishListByid(id string) (*models.Wishlist, error)
	GetWishListsByUserId(userId string) ([]models.Wishlist, error)
	UpdateWishList(id string, wishlist *models.Wishlist) error
	DeleteWishList(id string) error
}

type wishListRepositoryImpl struct{}

func NewWishListRepository() WishListRepository {
	return &wishListRepositoryImpl{}
}

func (w *wishListRepositoryImpl) CreateWishList(wishlist *models.Wishlist) error {
	result := config.DB.Create(wishlist)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (w *wishListRepositoryImpl) DeleteWishList(id string) error {
	// Infere a tabela a ser utilizada para efetuar a operação com base no tipo passado dentro do Delete

	// Passar &models.Wishlist garante que apenas o filtro no Where seja aplicado. Dessa forma, passar u struct vazio garante que apenas o ID seja usado para exclusão, pois assi o GORM ignora os demais campos do modelo, por exemplo, excluir alem da condicao do ID tambem ir pelo nome, ou email. 

	// Se eu tentar usar Delete(wishlist), o GORM tentaria excluir com base e todos os campos preenchidos do objeto wishlist, levando a comportamentos indesejados.
	result := config.DB.Where("id = ?", id).Delete(&models.Wishlist{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("wishlist not found")
	}

	return nil
}

func (w *wishListRepositoryImpl) GetWishListByid(id string) (*models.Wishlist, error) {
	// o GORM precisa passar o ponteiro da struct como parametro no metodo First. pois ele precisa de identificar sobre qual modelo (tabela) essa operação será feita, alem de armazenar essa consulta dentro da struct para poder retornar esses dados.
	var wishlist models.Wishlist
	result := config.DB.First(&wishlist, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return &wishlist, nil
}

func (w *wishListRepositoryImpl) GetWishListsByUserId(userId string) ([]models.Wishlist, error) {
	var wishLists []models.Wishlist
	result := config.DB.Where("user_id = ?", userId).Find(&wishLists)
	if result.Error != nil {
		return nil, result.Error
	}
	return wishLists, nil
}

func (w *wishListRepositoryImpl) UpdateWishList(id string, wishlist *models.Wishlist) error {
	// equivalente ao dicionario no Python e C#
	updateData := map[string]interface{}{}

	if wishlist.Title != "" {
		updateData["title"] = wishlist.Title
	}

	if wishlist.Description != "" {
		updateData["description"] = wishlist.Description
	}

	if !wishlist.EventDate.IsZero() {
		updateData["event_date"] = wishlist.EventDate
	}

	result := config.DB.Model(&models.Wishlist{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("wishlist not found or no changes detected")
	}

	return nil
}


