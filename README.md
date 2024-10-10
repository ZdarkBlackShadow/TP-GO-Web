### TP-GO-Web

## Description
 L'objectif de ce TP est de mettre en pratique les notions théoriques vues sur la programmation avec Golang orientée Web. Ce TP comporte 3 challenges, un où je dois implémenter une route **/promo** qui permet d'afficher les informations liées à une classe ainsi que la liste des étudiants qui la composent, un où je dois créer une route **/change** qui affiche une page avec un message différent en fonction du numéro de vues de la page (pair ou impair). Le compteur démarre à 0 au lancement du serveur et s'incrémente chaque fois que l'utilisateur accède à cette route, un où je dois créer une route **/user/form** qui contient un formulaire permettant à l'utilisateur de renseigner ses informations personnelles et une dernière route **/user/display** qui affichera les données renseignées par l'utilisateur : nom,prénom, date de naissance, et sexe. Si des données sont manquantes, ou bien non renseignées affichez un message demandant à l’utilisateur de renseigner ses
informations personnelles. Il y a également une route **/user/treatment** pour traiter les données soumises via le formulaire et qui redirige vers **/user/display**.

## Requirements : 
- Avoir Git ([Télécharger Git](https://git-scm.com/downloads))
- Avoir goland ([Télécharger Goland](https://go.dev/dl))
- Compatible avec Linux, MacOS, Chrome OS et Windows

## Instructions
- **Instalation**
    - 1 : Aller dans le ficher où vous vouler stocker le projet
    - 2 : Ouvrir un terminal dans ce ficher(ouvrez votre explorateur de ficher, ouvrer le ficher où vous voulez stocker le TP, click droit sur la souris, Plus d'options, Ouvrir dans le terminal)
    - 3 : Taper **_git clone https://github.com/ZdarkBlackShadow/TP-GO-Web.git_** et appuyer sur la touche **_entrer_**
- **Démarrage**
    - 1 : Ouvrir un terminal dans ce ficher(ouvrez votre explorateur de ficher, ouvrer le ficher où vous voulez stocker le TP, click droit sur la souris, Plus d'options, Ouvrir dans le terminal)
    - 2 : Taper **_go run main.go_** et appuyer sur la touche **_entrer_**
    - 3 : [Clickez ici pour regarder le TP](http://localhost:8000/promo)

## Auteur
- **Adrien Lecomte** [Mon github](https://github.com/ZdarkBlackShadow)