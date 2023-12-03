document.addEventListener('DOMContentLoaded', () => {
    const apiUrl = 'http://localhost:8080/f1teams/';
    const apiUrl2 = 'http://localhost:8080/f1teams';
    const resultContainer = document.getElementById('result-container');
    const teamIdInput = document.getElementById('teamId');
    const teamNameInput = document.getElementById('teamName');
    const driver1Input = document.getElementById('driver1');
    const driver2Input = document.getElementById('driver2');
    const carInput=document.getElementById('car');
    const popintsInput = document.getElementById('points');
    const wcInput = document.getElementById('wc');
    const clasificationInput = document.getElementById('clasification');
    
  
    const searchButton = document.getElementById('getTeam');
    const getAllTeamsButton = document.getElementById('getAllTeams');
    const createTeamButton = document.getElementById('createTeam');
    const updateTeamButton = document.getElementById('updateTeam');
    const deleteTeamButton = document.getElementById('deleteTeam');
  
    let currentTEAM = null;

    searchButton.addEventListener('click', () => {
      const teamId = teamIdInput.value;

      console.log(teamIdInput.value);
      fetch(`${apiUrl}${teamId}`)
        .then((response) => response.json())
        
        .then((data) => {
          currentTEAM = data;
          document.getElementById("teamId2").textContent = currentTEAM.id;
          document.getElementById("teamEquipo").textContent = currentTEAM.equipo;
          document.getElementById("teamDriver1").textContent = currentTEAM.driver_1;
          document.getElementById("teamDriver2").textContent = currentTEAM.driver_2;
          document.getElementById("teamCarro").textContent = currentTEAM.carro;
          document.getElementById("teamPuntos").textContent = currentTEAM.puntos;
          document.getElementById("teamCampeonatoConstructores").textContent = currentTEAM.campeonatoconstructores;
          document.getElementById("teamClasificacion").textContent = currentTEAM.clasificacion;

// Mostrar los elementos
          document.getElementById("teamInfo").classList.remove("hidden");
        })
        .catch((error) => {
          
          //displayError('Error al realizar la solicitud.');
        });
    });
  
    getAllTeamsButton.addEventListener('click', () => {
      fetch(apiUrl2)
        .then((response) => response.json())
        .then((data) => {
            currentTEAM = data; // Almacenar el Pokémon actual en la variable global
            console.log(data) 
            
            // Convertir los objetos a una cadena legible
            const formattedData = JSON.stringify(data, null, 2)
              .replace(/"/g, '')  
              .replace(/{/g, '') 
              .replace(/,/g, '')  
              .replace(/{/g, '')  
              .replace(/\[/g, '')
              .replace(/\]/g, '')  
              .replace(/}/g, '');  


            // Mostrar la cadena en la consola para verificar
            console.log(formattedData);

            // Asignar la cadena al textarea
            document.getElementById("teamInfo2").value = formattedData;


            // Mostrar el textarea
            


// Mostrar los elementos
          document.getElementById("teamInfo").classList.remove("hidden");
        })
        .catch((error) => {
          displayError('Error al realizar la solicitud.');
          console.log(error);

        });
    });
  
    createTeamButton.addEventListener('click', () => {
      const teamData = {
          equipo: teamNameInput.value,
          driver_1: driver1Input.value,
          driver_2: driver2Input.value,
          carro: carInput.value,
          puntos: parseInt(popintsInput.value, 10),  
          campeonato: wcInput.value,   
          clasificacion: parseInt(clasificationInput.value, 10)
      };
  
      fetch(apiUrl2, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify(teamData),
      })
          .then((response) => response.json())
          .then((data) => {
              displayResult(data);
          })
          .catch((error) => {
              displayError('Error al realizar la solicitud.');
          });
  });
  
    updateTeamButton.addEventListener('click', () => {
      const teamId = teamIdInput.value;
      const teamData = {
        equipo: teamNameInput.value,
        driver_1: driver1Input.value,
        driver_2: driver2Input.value,
        carro: carInput.value,
        puntos: parseInt(popintsInput.value, 10),  
        campeonatoconstructores: wcInput.value,   
        clasificacion: parseInt(clasificationInput.value, 10)
      };
  
      fetch(`${apiUrl}${teamId}`,{
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(teamData),
      })
        .then((response) => response.json())
        .then((data) => {
          displayResult(data);
        })
        .catch((error) => {
          displayError('Error al realizar la solicitud.');
        });
    });
  
    deleteTeamButton.addEventListener('click', () => {
      const teamId = teamIdInput.value;
  
      fetch(`${apiUrl}${teamId}`, {
        method: 'DELETE',
      })
        .then((response) => response.json())
        .then((data) => {
          displayResult(data);
        })
        .catch((error) => {
          displayError('Error al realizar la solicitud.');
        });
    });
  
    function displayResult(data) {
      resultContainer.innerHTML = JSON.stringify(data, null, 2);
    }
  
    function displayError(message) {
      resultContainer.innerHTML = `<p class="error">${message}</p>`;
    }
  });
  