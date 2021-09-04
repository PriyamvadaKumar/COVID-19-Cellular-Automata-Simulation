# COVID-19-Cellular-Automata-Simulation
# TABLE OF CONTENTS

- AIM

- TECHNOLOGIES

- SETUP

- PREREQUISITE/PACKAGES IMPORTED

- TERMINAL INPUT/USAGE

- STATUS

- SOURCES


# AIM

The aim of the project was to use a Cellular automata based model to simulate the spread of Coronavirus and analyze the results of infection in terms of the number of people infected, recovered, dead and susceptible in each of the 4 modeling cases namely - Moore neighborhood +basic transition, Moore neighborhood + random sampling transition
Random neighborhood + basic transition and Random neighborhood + random sampling transition.

Along with this, we also see the effects of incorporating a recovery period, vaccination rate, quarantine probability , and case fatality rate parameters.

Modeled the spread of disease in terms of susceptible, infected and recovered individuals using a 1D CA board based on  https://link.springer.com/article/10.1007/s11430-009-0044-9 - using time periods and rate of infection , recovery and cure.


# TECHNOLOGIES
- JUPYTER NOTEBOOK
- PYTHON 3.7.6
- Atom Text Editor
- Go Language 


# SETUP
Open readme.md using ATOM
Place the code folders under Src .
**Canvas, gifhelper packages must also be there .
In terminal , enter the directory/folder say COVID/COVID1. Do go build


# PREREQUISITE / PACKAGES IMPORTED

Imported:
- reflect
- testing
- Canvas
- Images
- fmt
- math/rand
- time
- gifhelper
- os
- strconv
- bufio
- strings

# TERMINAL INPUT/USAGE

# For covid folder, the command line is of the form:

./covid readparam/readdata filename twodigitcode numbdays rv vr qp cfr

where, parameters are :

- covid: command os.Args[0]
- readparam or readdata :Os.Arg[1] , parameter that decides whether to generate data for initial board randomly by reading parameters from a file  or by reading persons status from a file
- filename os.Args[2], name of the file - initial_data.txt or parameter_data.txt
- numbdays= number of days the simulation is run =100
- rv = recovery period = 15
- vr = vaccination rate (% per day)= 1
- qp = quarantine probability =0.5
- cfr = case fatality rate (% per day)=0.5
- twodigitcode refers to neighborhood choice+transition type

Here, 00-Moore neighborhood + basic transition

01- Moore neighborhood + random sampling transition

10- Random neighbourhood+ basic transition

11- random neighborhood +random sampling transition



May change number of days of simulation , initial_data.txt and parameter_data.txt file , rv,vr,qp, cfr to get different results and graphs (Currently assigned arbitrarily )



Graphs for number of infected/recovered/susceptible/vaccinated/dead individuals vs number of days outputted from Jupyter Notebook 

Example output : 

<img width="558" alt="Screen Shot 2021-09-04 at 1 44 23 PM" src="https://user-images.githubusercontent.com/77410526/132103703-094ac2b1-a52a-4189-ae08-8baf0c2584ce.png">

        x axis : Number of days 

Here, the following parameters have been set - days of running simulation = 100 , recovery period = 15,  vaccination rate(% per day)=1, quarantine probability=0.5, ,case fatality rate = 0.5 


First 4 commands to be entered in terminal the following for when initial_data.txt file used as input(manually creating board with susceptible and infected)

- ./covid readdata initial_data.txt 00 100 15 1 0.5 0.5
- ./covid readdata initial_data.txt 01 100 15 1 0.5 0.5
- ./covid readdata initial_data.txt 10 100 15 1 0.5 0.5
- ./covid readdata initial_data.txt 11 100 15 1 0.5 0.5




Next 4 terminal commands are for when board is assigned randomly with susceptible and infected using parameter_data.txt files

- ./covid readparam parameter_data.txt 00 100 15 1 0.5 0.5
- ./covid readparam parameter_data.txt 01 100 15 1 0.5 0.5
- ./covid readparam parameter_data.txt 10 100 15 1 0.5 0.5
- ./covid readparam parameter_data.txt 11 100 15 1 0.5 0.5




# For covid1 folder ,

Go inside covid1 foder, do go build and then enter ./covid1 in terminal




# STATUS

The project is complete but can be expanded to take real COVID data as input and compare the results of the CA based model with Real life COVID spread data.
The model may further be improved by introducing more parameters to encompass all the possible states like individual age, type of treatment , spontaneous infection rate(when source of infection cannot be traced), Immigration/emigration rate (Area model is open or closed , influenced by spontaneous infection rate) etc
The model maybe be futher improved to have a stronger statistical basis




# SOURCES

1.    White S H, del Rey A M, Sánchez G R. Modeling epidemics using cellular automata. Appl Math Comput, 2007, 186: 193―202
2.    Mikler A R, Venkatachalam S, Abbas K. Modeling infectious diseases using global stochastic cellular automata. J Biol Syst, 2005, 13:421―439
3.    Wang J F, McMichael A J, Meng B, et al. Spatial dynamics of an epidemic of severe acute respiratory syndrome in an urban area. Bull World Health Organ, 2006,   84: 965―968
4. https://www.sciencedirect.com/topics/computer-science/cellular-automata
5. https://plato.stanford.edu/entries/cellular-automata/
6.    Bombardt J N. Congruent epidemic models for unstructured and structured populations: Analytical reconstruction of a 2003 SARS outbreak. Math Biosci, 2006, 203: 171―203

7. https://link.springer.com/article/10.1007/s11430-009-0044-9
8. https://www.mivegec.ird.fr/images/stories/PDF_files/0507.pdf
9.  Xiao X, Shao SH, Chou KC. A probability cellular automaton model for hepatitis b viral infections. Biochem Biophys Res Commun. 2006;342:605–15. [PubMed] [Google Scholar]
10. https://en.wikipedia.org/wiki/Mathematical_modelling_of_infectious_disease
11. https://www.cdc.gov/coronavirus/2019-ncov/index.html
12. https://towardsdatascience.com/algorithmic-beauty-an-introduction-to-cellular-automata-f53179b3cf8f
