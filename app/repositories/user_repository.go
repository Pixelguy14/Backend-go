package repositories

type UserRepository struct  {
	Client *firestore.Client
}

func (r *UserRepository) CreateUser (user models.Usuario) error {
	_, _, err := r.Client.Collection("usuario").Add(context.Background(),user)
	return err 
}