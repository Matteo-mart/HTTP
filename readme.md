## Présentation du projet

Ce projet consiste en la création d'une API REST robuste développée en Go. Elle permet de gérer une base de données d'albums musicaux. L'application expose des endpoints HTTP pour consulter, ajouter ou modifier des enregistrements stockés dans une base de données SQL.

L'accent a été mis sur la séparation des responsabilités (architecture modulaire) et sur la qualité logicielle via la mise en place de tests unitaires.
## Fonctionnalités Techniques

    CRUD Complet : Gestion des albums (Création, Lecture, Mise à jour).

    Persistance SQL : Connexion et requêtes sécurisées vers une base de données.

    Architecture Modulaire :

        recordings/ : Logique métier et accès aux données.

        tests/ : Suite de tests automatisés pour valider le bon fonctionnement de l'API.

    Sécurité des accès : Utilisation de variables d'environnement pour les identifiants de connexion.

## Compétences SIO (SLAM) validées

    B1.1 - Gérer le patrimoine informatique : Utilisation de tests unitaires pour assurer la non-régression du code.

    B1.3 - Développer la présence en ligne : Conception et réalisation d'un service web (Back-end).

    Gestion des données : Manipulation d'une base de données relationnelle via le langage Go.

## Installation et Lancement
### 1. Configuration de l'environnement

    Le projet utilise des variables d'environnement pour sécuriser l'accès à la base de données.
### 2. Exécution du serveur

    Dans un terminal, placez-vous à la racine du projet et lancez la commande suivante :
        Bash

        DBUSER=votre_utilisateur DBPASS=votre_mot_de_passe go run .

### 3. Exécution des tests

    Pour vérifier que toutes les fonctions (comme getAlbum) fonctionnent correctement avec la base de données :
        Bash

        go test ./tests

## Structure du dépôt
    Fichier / Dossier	Rôle
        main.go	Initialisation du serveur et définition des routes HTTP.
/recordings	Contient les méthodes d'interaction avec la base de données.
/tests	Regroupe l'ensemble des scripts de tests du package.
go.mod	Gestionnaire des modules et dépendances Go.
