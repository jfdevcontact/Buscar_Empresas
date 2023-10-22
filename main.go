package main

import(
	"log"
	"net/http"
	"encoding/json"
	"io"
    "fmt"
	"time"
	
	
	
	
	
)

func main(){

	type Get_cnpj struct {
		ID                    string
		Status                string                  `json:"status"`
		UltimaAtualizacao     time.Time               `json:"ultima_atualizacao"`
		Cnpj                  string                  `json:"cnpj"`
		Tipo                  string                  `json:"tipo"`
		Porte                 string                  `json:"porte"`
		Nome                  string                  `json:"nome"`
		Fantasia              string                  `json:"fantasia"`
		Abertura              string                  `json:"abertura"`
		AtividadePrincipal    string                  `json:"atividade_principal"`
		AtividadesSecundarias string                  `json:"atividades_secundarias"`
		NaturezaJuridica      string                  `json:"natureza_juridica"`
		Logradouro            string                  `json:"logradouro"`
		Numero                string                  `json:"numero"`
		Complemento           string                  `json:"complemento"`
		Cep                   string                  `json:"cep"`
		Bairro                string                  `json:"bairro"`
		Municipio             string                  `json:"municipio"`
		Uf                    string                  `json:"uf"`
		Email                 string                  `json:"email"`
		Telefone              string                  `json:"telefone"`
		Efr                   string                  `json:"efr"`
		Situacao              string                  `json:"situacao"`
		DataSituacao          string                  `json:"data_situacao"`
		MotivoSituacao        string                  `json:"motivo_situacao"`
		SituacaoEspecial      string                  `json:"situacao_especial"`
		DataSituacaoEspecial  string                  `json:"data_situacao_especial"`
		CapitalSocial         string                  `json:"capital_social"`
		
	}
	


	var Input_cnpj string

	fmt.Println("Digite o cnpj que sera consultado EX: 00000000000000")
	fmt.Scan(&Input_cnpj)

	fmt.Println()

	Req, err := http.Get("https://receitaws.com.br/v1/cnpj/" + Input_cnpj + "/" )
	if err != nil{
		log.Println(err)
	}

	Body, err := io.ReadAll(Req.Body)
	if err != nil{
		log.Println(err)
	}

	var Get1 Get_cnpj
	
	json.Unmarshal(Body, &Get1)
	
	
	fmt.Println("\033[32mAbertura: ", Get1.Abertura)
	fmt.Println("\033[32mSituação:", Get1.Situacao)
	fmt.Println("\033[32mTipo: ", Get1.Tipo)
	fmt.Println("\033[32mNome: ", Get1.Nome)
	fmt.Println("\033[32mNome fantasia: ", Get1.Fantasia)
	fmt.Println("\033[32mPorte da empresa: ", Get1.Porte)
	fmt.Println("\033[32mNatureza juridica: ", Get1.NaturezaJuridica)
	fmt.Println("\033[32mAtividade principal:", Get1.AtividadePrincipal)

	fmt.Println()

	fmt.Println("\033[32mInfomações do Endereço da empresa: ", Get1.Fantasia)
	fmt.Println()

	fmt.Println("\033[32mRua: ", Get1.Logradouro)
	fmt.Println("\033[32mNumero: ", Get1.Numero)
	fmt.Println("\033[32mComplemento: ", Get1.Complemento)
	fmt.Println("\033[32mMunicipio: ", Get1.Municipio)
	fmt.Println("\033[32mBairro: ", Get1.Bairro)
	fmt.Println("\033[32mEstado: ", Get1.Uf)
	fmt.Println("\033[32mCep: ", Get1.Cep)

	fmt.Println()
	fmt.Println("\033[32mInfoamaçõs de contato: ")
	fmt.Println()

	fmt.Println("\033[32mTelefone:", Get1.Telefone)
	fmt.Println("\033[32mEmail: ", Get1.Email)

	fmt.Println()

	fmt.Println("\033[32mData situação", Get1.DataSituacao)
	fmt.Println("\033[32mMotivo situação: ", Get1.MotivoSituacao)
	fmt.Println("\033[32mSituação especial: ", Get1.DataSituacaoEspecial)
	fmt.Println("\033[32mCapital social: ", Get1.CapitalSocial)



}

	

