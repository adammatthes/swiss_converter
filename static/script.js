'use strict';

async function getConversionOptions(startingType, category) {
	const url = '/api/get-conversion-options';
	const data = {value: 0, type: startingType, category: category};

	const response = await fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(data),
	});

	if (!response.ok) {
		throw new Error(`getConversionOptions failed: ${response.status}`);
	}

	const result = await response.json();

	if (!result.options) {
		return ["No conversion types found"];
	}

	return result.options;
}

async function getStartingTypes(category) {
	const url = '/api/get-starting-types';
	const data = {value: 0, type: category};

	const response = await fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(data),
	});

	if (!response.ok) {
		throw new Error(`getStartingTypes failed: ${response.status}`);
	}

	const result = await response.json();

	if (!result.options) {
		return ["No starting types found"];
	}

	return result.options;
}

async function getConversionResult() {
	const categoryDrop = document.getElementById('categorySelect');
	const category = categoryDrop.value;

	const startType = document.getElementById('startingTypeSelect');
	const start = startType.value;

	const endType = document.getElementById('destinationTypeSelect');
	const end = endType.value;

	const input = document.getElementById('userInput');
	const value = input.value;

	const url = category === 'Currency' ? '/api/currency': '/api/convert';
	const data = {
		"category": category,
		"start-type": start,
		"end-type": end,
		"value": value,
	}

	const response = await fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(data),
	});

	if (!response.ok) {
		const errorMessage = await response.text();

		throw new Error(`getConversionResult failed: ${errorMessage}`);
	}

	const result = await response.json();

	if (!result.result) {
		return "Invalid Input"
	}

	return result.result
}

async function createNewConversion() {
	const startInput = document.getElementById("customStartInput");
	const endInput = document.getElementById("customEndInput");
	const exchangeInput = document.getElementById("customExchangeRateInput")

	const data = {
		"start-type": startInput.value,
		"end-type": endInput.value,
		"value": exchangeInput.value,
	}
	const url = "api/create-conversion"

	const response = await fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(data),
	});

	if (!response.ok) {
		const errorMessage = await response.text();
		throw new Error(`createNewConversion failed: ${errorMessage}`);
	}

	const result = await response.json()

	console.log(result)
}

function makeSelectElement(options, id) {
	if (!options || options.length === 0) {
		return
	}

	const select = document.createElement('select');
	select.className = 'dropdownMenu';
	select.id = id;

	const disabledOption = document.createElement('option');
	disabledOption.disabled = true;
	disabledOption.selected = true;
	disabledOption.innerHTML = "Select an option";
	select.appendChild(disabledOption);

	for (let opt of options) {
		const newOption = document.createElement('option');
		newOption.value = opt;
		newOption.innerHTML = opt;
		select.appendChild(newOption);
	}

	return select;
}

function modifySelectOptions(targetSelect, newOptions) {
	const newChildren = [];
	const disabledOption = document.createElement('option');
	disabledOption.disabled = true;
	disabledOption.selected = true;
	disabledOption.innerHTML = "Select an option";
	newChildren.push(disabledOption);
	for (let opt of newOptions) {
		const newOpt = document.createElement('option');
		newOpt.value = opt;
		newOpt.innerHTML = opt;
		newChildren.push(newOpt);
	}
	targetSelect.replaceChildren(...newChildren);
}

async function generateDestinationTypeSelect() {

	const selectedValue = document.getElementById('startingTypeSelect').value;
	const category = document.getElementById('categorySelect').value;

	const options = await getConversionOptions(selectedValue, category);

	let destSelect = document.getElementById('destinationTypeSelect');
	if (!destSelect) {
		destSelect = makeSelectElement(options, 'destinationTypeSelect');
	} else {
		modifySelectOptions(destSelect, options);
		return;
	}

	destSelect.addEventListener('change', async function() {
		let input = createInputField();
		if (input) {
			const menu = document.getElementById('conversionMenu');
			menu.appendChild(input);
			return;
		}
		input = document.getElementById('userInput');
		const dest = document.getElementById('destinationTypeSelect');
		const start = document.getElementById('startingTypeSelect');
		if (input.value && start.value != 'Select an option' && dest.value != 'Select an option') {
			await submitInput();
		}
	});

	const menu = document.getElementById('conversionMenu');
	menu.appendChild(destSelect);

}

async function submitInput() {
	const start = document.getElementById('startingTypeSelect');
	const dest = document.getElementById('destinationTypeSelect');
	const ipt = document.getElementById('userInput');
	if (start.value === 'Select an option' || dest.value === 'Select an option' || !ipt.value) {
		return;
	}

	const result = await getConversionResult();

	let display = document.getElementById('resultOutput');
	if (!display) {
		display = document.createElement('p');
		display.id = 'resultOutput';

		const menu = document.getElementById('conversionMenu');
		menu.appendChild(display);
	}
	display.innerHTML = '';
	display.innerHTML = result;
}


function createInputField() {
	let input = document.getElementById('userInput');
	if (input) {
		return;
	}

	input = document.createElement('input');
	input.id = 'userInput';
	input.type = 'text';
	input.className = 'inputField';

	input.addEventListener('input', submitInput);
	return input;
}

async function generateStartingTypeSelect() {
	const selectedValue = document.getElementById('categorySelect').value;
	const options = await getStartingTypes(selectedValue);

	let startSelect = document.getElementById('startingTypeSelect');
	if (!startSelect) {
		startSelect = makeSelectElement(options, 'startingTypeSelect');
		startSelect.addEventListener('change', async function() { generateDestinationTypeSelect();});
	} else {
		modifySelectOptions(startSelect, options);
		await generateDestinationTypeSelect();
		return;
	}

	const menu = document.getElementById('conversionMenu');
	menu.appendChild(startSelect);
}

function generateCustomGenerationFields() {
	if (document.getElementById('customGenerationSubmitButton')) {
		return;
	}

	const customFields = document.getElementById("customRateFields");
	
	const firstInput = document.createElement("input");
	firstInput.type = 'text';
	firstInput.id = 'customStartInput';
	const secondInput = document.createElement("input");
	secondInput.type = 'text';
	secondInput.id = 'customEndInput';
	const thirdInput = document.createElement("input");
	thirdInput.type = 'text';
	thirdInput.id = 'customExchangeRateInput';
	const submitGeneration = document.createElement("button");
	submitGeneration.id = 'customGenerationSubmitButton';
	submitGeneration.textContent = 'Generate Conversion';
	submitGeneration.addEventListener('click', async function() { await createNewConversion();});



	customFields.appendChild(firstInput);
	customFields.appendChild(secondInput);
	customFields.appendChild(thirdInput);
	customFields.appendChild(submitGeneration);
}

const firstDropDown = document.getElementById('categorySelect');
firstDropDown.addEventListener('change', async function() { generateStartingTypeSelect();});

const addButton = document.getElementById('customRateInitiator');
addButton.addEventListener('click', function() { generateCustomGenerationFields();});
