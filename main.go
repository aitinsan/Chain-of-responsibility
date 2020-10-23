package main

import "fmt"

type People struct {
	Money, Age int
	Name, Job string
	HealthGood bool
}
func (p *People) String() string{
	return fmt.Sprintf("Name: %s Jod: %s Age: %d Money: %d HealthGood: %t",p.Name, p.Job, p.Age, p.Money,p.HealthGood)
}
func NewPeople(name,job string, age,money int , health bool) *People{
	return &People{money,age, name, job, health}
}
type BadHabit interface{
	Add(b BadHabit)
	Apply()
}
type BadHabitOfPerson struct{
	next BadHabit
	people *People
}
func (b *BadHabitOfPerson) Add(bad BadHabit){
	if b.next != nil{
		b.next.Add(bad)
	} else {
		b.next = bad
	}
}
func (b *BadHabitOfPerson) Apply(){
	if b.next != nil {
		b.next.Apply()
	}
}
func NewBadHabit(p *People) *BadHabitOfPerson {
	return &BadHabitOfPerson{people: p}
}
type Sigaret struct{
	BadHabitOfPerson
}
func (s *Sigaret) Apply(){


	if s.people.Money <10{
		fmt.Printf("Money is not enought for sigaret \n")
	}
	if s.people.Money >=10{
		fmt.Printf("Smoked sigaret,%s \n",s.people.Name)
		if s.people.HealthGood == true{
			s.people.HealthGood=false
		}
		s.people.Money-=10
	}
	s.BadHabitOfPerson.Apply()
}
func NewSigarets(p *People) *Sigaret{
	return &Sigaret{BadHabitOfPerson{people: p}}
}
type Iqos struct{
	BadHabitOfPerson
}
func (i *Iqos) Apply(){


	if i.people.Money <20{
		fmt.Printf("Money is not enought for Iqos\n")
	}
	if i.people.Money >=20{
		i.people.Money-=20
		fmt.Printf("Smoked Iqos ,%s\n",i.people.Name)
		if i.people.HealthGood == true{
			i.people.HealthGood=false
		}
	}
	i.BadHabitOfPerson.Apply()
}
func NewIqosSmoke(p *People) *Iqos{
	return &Iqos{BadHabitOfPerson{people: p}}
}
type Dota struct{
	BadHabitOfPerson
}
func (d *Dota) Apply(){
	if d.people.Money <30{
		fmt.Printf("Money is not enought for Dota\n")
	}
	if d.people.Money >=30{
		d.people.Money-=30
		fmt.Printf("Played Dota ,%s\n",d.people.Name)
		if d.people.HealthGood == false{
			d.people.HealthGood=true
		}
	}

	d.BadHabitOfPerson.Apply()

}
func NewDotaGame(p *People) *Dota{
	return &Dota{BadHabitOfPerson{people: p}}
}
func main(){
	Tleuzhan :=NewPeople("Tleuzhan","Professor", 6, 60,true)
	fmt.Println(Tleuzhan)
	Habbit:=NewBadHabit(Tleuzhan)
	Habbit.Add(NewSigarets(Tleuzhan))
	Habbit.Add(NewIqosSmoke(Tleuzhan))
	Habbit.Add(NewDotaGame(Tleuzhan))
	Habbit.Apply()
	fmt.Println(Tleuzhan)
}
