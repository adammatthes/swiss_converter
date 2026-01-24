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

	if (value === '-') {
		return '0';
	}

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

	if (startInput.value === '' || endInput.value === '' || exchangeInput.value === '') {
		return;
	}

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

	const createSection = document.getElementById("customGenerationSection");

	createFadingParagraph(result.message, "#50C878", createSection);

	setTimeout(() => {document.location.reload(true)}, 5000);
}

async function deleteCustomConversion() {
	const startInput = document.getElementById("deleteStartInput");
	const endInput = document.getElementById("deleteEndInput");

	if (startInput.value === '' || endInput.value === '') {
		return;
	}

	const data = {
		'start-type': startInput.value,
		'end-type': endInput.value,
		'value': null,
	}
	const url = "api/delete-conversion"

	const response = await fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(data),
	});

	if (!response.ok) {
		const errorMessage = await response.text();
		throw new Error(`deleteCustomConversion failed: ${errorMessage}`);
	}

	const result = await response.json()

	const deleteSection = document.getElementById('deleteConversionSection');

	createFadingParagraph(result.message, '#50C878', deleteSection);

	setTimeout(() => {document.location.reload(true)}, 5000);
}

async function updateCurrencyRates() {
	const url = "api/update-currencies";

	const response = await fetch(url);
	if (!response.ok) {
		throw new Error(`updateCurrencyRates Failed: ${response.status}`)
	}

	const result = await response.json()

	const menu = document.getElementById('conversionMenu')

	createFadingParagraph(JSON.stringify(result), '#50C878', menu)
}

function createFadingParagraph(text, color, attachTo) {
	const p = document.createElement('p');
	p.textContent = text;
	p.style.background = color;

	attachTo.appendChild(p);

	setTimeout(() => {
		p.classList.add('fade-out');
	}, 2000);

	setTimeout(() => {
		p.remove();
	}, 3000);

}

async function getMetricsData() {
	const url = "api/metrics";

	const response = await fetch(url);
	if (!response.ok) {
		const errorMessage = await response.text();
		throw new Error(`getMetricsData failed: ${errorMessage}`);
	}

	const result = await response.json();

	return result;
}

async function makeMetricsTable() {
	const data = await getMetricsData();
	if (!data) {
		return;
	}

	const colNames = ["Start Type", "End Type", "Total Number of Conversions"];

	let metricsTable = document.getElementById('metricsTable');
	if (metricsTable) {
		metricsTable.replaceChildren();
	} else {
		metricsTable = document.createElement('table');
		metricsTable.id = 'metricsTable';
	}

	const metricsTableHead = document.createElement('thead');
	metricsTableHead.className = "tableHead";

	for (const col of colNames) {
		const nextCol = document.createElement("th");
		nextCol.value = col;
		nextCol.innerHTML = col;
		nextCol.className = 'columnHeader';
		metricsTableHead.appendChild(nextCol);
	}

	metricsTable.style.border = "1px solid black";
	metricsTable.style.borderRadius = "8px";
	metricsTable.style.boxShadow = "1px 1px black";

	metricsTable.appendChild(metricsTableHead);

	const tBody = document.createElement('tbody');

	for (const row of data) {
		const nextRow = document.createElement('tr');
		const nextStartType = document.createElement('td');
		nextStartType.innerHTML = row.StartType;
		nextRow.appendChild(nextStartType);

		const nextEndType = document.createElement('td');
		nextEndType.innerHTML = row.EndType;
		nextRow.appendChild(nextEndType);

		const nextCount = document.createElement('td');
		nextCount.innerHTML = row.Count;
		nextRow.appendChild(nextCount)

		tBody.appendChild(nextRow)
	}

	metricsTable.appendChild(tBody);

	const metricsSection = document.getElementById('metricsSection');
	metricsSection.appendChild(metricsTable);
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
		let input = createInputField('userInput');
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

		const outputSection = document.getElementById('resultOutputSection');
		outputSection.appendChild(display);
	}
	display.innerHTML = '';
	display.innerHTML = result;
}


function createInputField(id) {
	let input = document.getElementById(id);
	if (input) {
		return;
	}

	input = document.createElement('input');
	input.id = id;
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
		startSelect.addEventListener('change', async function() { await generateDestinationTypeSelect();});
	} else {
		modifySelectOptions(startSelect, options);
		await generateDestinationTypeSelect();
		return;
	}

	const menu = document.getElementById('conversionMenu');
	menu.appendChild(startSelect);

	if (selectedValue === 'Currency') {
		const updateButton = document.createElement('button');
		updateButton.id = 'currencyUpdateButton';
		updateButton.textContent = 'Update Currencies';
		updateButton.addEventListener('click', async function() { await updateCurrencyRates();});
		menu.appendChild(updateButton);
	} else {
		const updateButton = document.getElementById('currencyUpdateButton');
		if (updateButton) {
			updateButton.remove();
		}
	}
}

function makeInputDiv(label, htmlFor) {
	const inputDiv = document.createElement('div');
	inputDiv.className = 'customField';

	const lbl = document.createElement('label');
	lbl.textContent = label;
	lbl.className = 'customCreateLabel';
	lbl.htmlFor = htmlFor;

	const textInput = document.createElement('input');
	textInput.type = 'text';
	textInput.id = htmlFor;
	textInput.className = 'inputField';

	inputDiv.appendChild(lbl);
	inputDiv.appendChild(textInput);

	return inputDiv;
}


function generateDeleteFields() {
	const deleteFields = document.getElementById('deleteCustomFields')

	const deleteStart = makeInputDiv('Delete Start Type', 'deleteStartInput');
	const deleteEnd = makeInputDiv('Delete End Type', 'deleteEndInput');

	deleteFields.appendChild(deleteStart);
	deleteFields.appendChild(deleteEnd);

	const submitDeleteButton = document.createElement('button');
	submitDeleteButton.id = 'submitDeleteButton';
	submitDeleteButton.textContent = 'Delete Conversion';
	submitDeleteButton.addEventListener('click', async function() { await deleteCustomConversion();});

	const deleteSection = document.getElementById('deleteConversionSection');

	deleteSection.appendChild(deleteFields);
	deleteSection.appendChild(submitDeleteButton);
}

function generateCustomGenerationFields() {
	if (document.getElementById('customGenerationSubmitButton')) {
		return;
	}

	const startLabel = document.createElement('label');
	startLabel.textContent = 'Start Type      ';
	startLabel.className = 'customCreateLabel';
	startLabel.htmlFor = 'customStartInput';

	const endLabel = document.createElement('label');
	endLabel.textContent = 'Destination Type';
	endLabel.className = 'customCreateLabel';
	endLabel.htmlFor = 'customEndInput';

	const rateLabel = document.createElement('label');
	rateLabel.textContent = 'Conversion Rate ';
	rateLabel.className = 'customCreateLabel';
	rateLabel.htmlFor = 'customExchangeRateInput';

	const customFields = document.getElementById("customRateFields");

	const startDiv = document.createElement('div');
	startDiv.className = 'customField';
	startDiv.appendChild(startLabel);
	const firstInput = createInputField("customStartInput");
	startDiv.appendChild(firstInput);

	const endDiv = document.createElement('div');
	endDiv.className = 'customField';
	endDiv.appendChild(endLabel);
	const secondInput = createInputField("customEndInput");
	endDiv.appendChild(secondInput);

	const rateDiv = document.createElement('div');
	rateDiv.className = 'customField';
	rateDiv.appendChild(rateLabel);
	const thirdInput = createInputField("customExchangeRateInput");
	rateDiv.appendChild(thirdInput);

	const submitGeneration = document.createElement("button");
	submitGeneration.id = 'customGenerationSubmitButton';
	submitGeneration.textContent = 'Generate Conversion';
	submitGeneration.addEventListener('click', async function() { await createNewConversion();});

	customFields.appendChild(startDiv);
	customFields.appendChild(endDiv);
	customFields.appendChild(rateDiv);

	const customSection = document.getElementById('customGenerationSection');
	customSection.appendChild(submitGeneration);
}

const firstDropDown = document.getElementById('categorySelect');
firstDropDown.addEventListener('change', async function() { generateStartingTypeSelect();});

const addButton = document.getElementById('customRateInitiator');
addButton.addEventListener('click', function() { generateCustomGenerationFields();});

const deleteButton = document.getElementById('deleteConversionButton');
deleteButton.addEventListener('click', function() { generateDeleteFields();});

const metricsButton = document.getElementById('metricsCreateButton');
metricsButton.addEventListener('click', async function() { makeMetricsTable();});
