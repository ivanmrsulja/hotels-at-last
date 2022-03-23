# Hotels At Last
HotelsAtLast je web aplikacija za rezervaciju hotelskih soba zasnovana na mikroservisnoj arhitekturi.

## Funkcionalnosti

### Neregistrovan korisnik
Može da vidi i pretražuje sve hotele u sistemu (pretraga se vrši prema cijeni od-do, broju ljudi koji mogu stati u sobu od-do, imenu i adresi) kao i da vidi njihovu ponudu (sobe). Pružena mu je mogućnost registacije kao i prijave na sistem.

### Registrovan korisnik
Nakon uspješne registracije i prijave, korisnik pored pregleda i pretrage hotela i pregleda njihovih ponuda (soba) ima mogućnost pregleda rezervacija te sobe kao i mogućnost rezervacije sobe na određen vremenski period ako je ona dostupna u istom.

### Admin
CRUD nad entitetima u sistemu - hotelima, sobama i korisnicima (za korisnike samo pregled i brisanje).

## Arhitektura sistema

API gateway - mikroservis sa kojim komunicira klijentska aplikacija i koji delegira zahtjeve ostalim mikroservisima. Tehnologija: Go

User auth service - mikroservis za autentifikaciju i autorizaciju korisnika. Tehnologije: Go, PostgreSQL baza

Hotel service - mikroservis koji je zadužen za CRUD operacije nad hotelima. Tehnologije: Go, PostgreSQL baza

Booking service - mikroservis koji je zadužen za rezervacije. Tehnologije: Rust, PostgreSQL baza

Klijentska aplikacija - klijentska aplikacija koja komunicira sa API gateway-om. Tehnologija: Vue.js
