# Calculate IR
![alt Script executed in terminal](https://s30.postimg.org/5notbsgld/Captura_de_tela_de_2017_01_03_21_55_04.png)

Projet to test [Salary](https://github.com/aalvesjr/salary) package

## Setup

```
mkdir -p $GOPATH/src/github.com/aalvesjr
cd $GOPATH/src/github.com/aalvesjr
git clone https://github.com/aalvesjr/calcula_ir.git
cd calcula_ir

glide install
go build
```

## Running

```
./calcula_ir 5168.78 127.79

#=> --------------Salary---------------
#=> Salary Gross         => R$ 5168.78
#=> Discounts            => R$ 127.79
#=> ---------------INSS----------------
#=> INSS Base            => R$ 5168.78
#=> INSS Rate            => 11.00%
#=> INSS Value           => R$ 568.57
#=> ----------------IR-----------------
#=> IR Base              => R$ 4600.21
#=> IR Rate              => 22.50%
#=> IR Without Discount  => R$ 1035.05
#=> IR Discount          => R$ 636.13
#=> IR Value             => R$ 398.92
#=> -----------------------------------
#=> Salary Net           => R$ 4073.51

```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request
