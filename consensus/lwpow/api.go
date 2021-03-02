type API struct {
	chain consensus.ChainReader

	lwpow *LWPoW
}

func (api *API) EchoNumber(ctx context.Context, number uint64) (uint64, error) {

	fmt.Println("called echo number")

	return number, nil

}