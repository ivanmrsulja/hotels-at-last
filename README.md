# Hotels At Last
HotelsAtLast je web aplikacija za rezervaciju hotelskih soba zasnovana na mikroservisnoj arhitekturi.

## Funkcionalnosti

### Neregistrovan korisnik
Može da vidi i pretražuje sve hotele u sistemu (pretraga se vrši prema cijeni od-do, broju ljudi koji mogu stati u sobu od-do, imenu, adresi, posedovanje parking mesta, klima uređaja, TV, mašina...) kao i da vidi njihovu ponudu (sobe). Pružena mu je mogućnost registacije kao i prijave na sistem.

### Registrovan korisnik
Nakon uspješne registracije i prijave, korisnik pored pregleda i pretrage hotela i pregleda njihovih ponuda (soba) ima mogućnost pregleda rezervacija te sobe kao i mogućnost rezervacije sobe na određen vremenski period ako je ona dostupna u istom. Korisniku je dostupna i opcija poništavanja rezervacije maksimalno 2 dana pre početka iste. Nakon provedenog boravka, korisnik može dati recenziju hotela uz ostavljanje komentara i ocjene od 1 do 5. Korisnik može prijaviti (njemu) nedolične komentare koji će biti poslati na uvid administratoru sistema.

### Admin
CRUD nad entitetima u sistemu - hotelima, sobama, korisnicima (za korisnike samo pregled, brisanje i banovanje na određeni period) i komentarima (pregled i brisanje prijavljenih komentara).

## Arhitektura sistema

API gateway - mikroservis sa kojim komunicira klijentska aplikacija i koji delegira zahtjeve ostalim mikroservisima. Tehnologija: Go

User auth service - mikroservis za autentifikaciju i autorizaciju korisnika. Tehnologije: Go, PostgreSQL baza

Hotel service - mikroservis koji je zadužen za CRUD operacije nad hotelima. Tehnologije: Go, PostgreSQL baza

Review service - mikroservis koji je zadužen za CRUD operacije nad recenzijama. Tehnologije: Go, PostgreSQL baza

Booking service - mikroservis koji je zadužen za rezervacije. Tehnologije: Rust, PostgreSQL baza

Klijentska aplikacija - klijentska aplikacija koja komunicira sa API gateway-om. Tehnologija: Vue.js
