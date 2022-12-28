package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrNotFound  = errors.New("cliente no existe")
	ErrExists    = errors.New("cliente ya existe")
	ErrDni       = errors.New("dni Nulo")
	ErrName      = errors.New("name Nulo")
	ErrDomicilio = errors.New("domicilio Nulo")
	ErrTelefono  = errors.New("telefono Nulo")
	ErrLegajo    = errors.New("legajo Nulo")
	ErrNulo      = errors.New("dato Nulo")
)

type Cliente struct {
	Legajo    string
	Nombre    string
	DNI       string
	Telefono  string
	Domicilio string
}

var clientes []Cliente

func (c Cliente) addCliente() (err error) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = ExistClienteByDNI(c.DNI)
	if err != nil {
		panic(err)
	}

	_, err = ValidarDato(c.DNI)
	if err != nil {
		panic(fmt.Errorf("%s %w", err, ErrDni))
	}
	_, err = ValidarDato(c.Nombre)
	if err != nil {
		panic(fmt.Errorf("%s %w", err, ErrName))
	}
	_, err = ValidarDato(c.Legajo)
	if err != nil {
		panic(fmt.Errorf("%s %w", err, ErrLegajo))
	}
	_, err = ValidarDato(c.Telefono)
	if err != nil {
		panic(fmt.Errorf("%s %w", err, ErrTelefono))
	}
	_, err = ValidarDato(c.Domicilio)
	if err != nil {
		panic(fmt.Errorf("%s %w", err, ErrDomicilio))
	}

	clientes = append(clientes, c)
	updateCliente(c)
	return
}

func verClientes() {
	for _, v := range clientes {
		fmt.Println(v)
	}
}

func ExistClienteByDNI(dni string) (b bool, err error) {
	for _, v := range clientes {
		if v.DNI == dni {
			err = ErrExists
			return
		}
	}
	return
}

func ValidarDato(campo string) (dato string, err error) {
	if campo != "" {
		return
	}
	err = ErrNulo
	return
}

func updateCliente(v Cliente) {
	text := fmt.Sprintf("\n%s,%s,%s,%s,%s", v.Legajo, v.Nombre, v.DNI, v.Telefono, v.Domicilio)

	f, err := os.OpenFile("customer.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func main() {

	defer func() {
		fmt.Println("ejecución finalizada")
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")

		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	//Carga de archivo
	file, err := os.Open("customer.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		cl := Cliente{items[0], items[1], items[2], items[3], items[4]}
		//fmt.Printf("Legajo: %s, Name: %s, DNI: %s, Telefono: %s, Direccion: %s \n", items[0], items[1], items[2], items[3], items[4])
		cl.addCliente()
	}
	//-----------------------------
	//Crear nuevo cliente
	cl := Cliente{"1010", "Raul", "1111", "12345678", "direccion1"}
	err = cl.addCliente()
	if err != nil {
		panic("ERROR CLIENTE NO AGREGADO")
	}
	verClientes()

	cl2 := Cliente{"1010", "", "1010", "12345678", "direccion1"}
	cl2.addCliente()
	if err != nil {
		panic("ERROR CLIENTE NO AGREGADO")
	}
	verClientes()
	/*
		for _, v := range clientes {
			line := fmt.Sprintf("%s,%s,%s,%s,%s", v.Legajo, v.Nombre, v.DNI, v.Telefono, v.Domicilio)
			file.WriteString(line)
		}*/

	/*cl3 := Cliente{"1010", "Raul", "1111", "12345678", "direccion1"}
	cl3.addCliente()
	cl3.verClientes()*/

}
