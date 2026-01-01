'use strict';

async function getConversionOptions(startingType) {
	const url = '/api/get-conversion-options';
	const data = {value: 0, type: startingType};

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

function makeSelectElement(options, id) {
	if (!options || options.length === 0) {
		return
	}

	const select = document.createElement('select');
	select.className = 'dropdownMenu';
	select.id = id;

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

	const options = await getConversionOptions(selectedValue);

	let destSelect = document.getElementById('destinationSelect');
	if (!destSelect) {
		destSelect = makeSelectElement(options, 'destinationSelect');
	} else {
		modifySelectOptions(destSelect, options);
	}

	const menu = document.getElementById('conversionMenu');
	menu.appendChild(destSelect);

}

const firstDropDown = document.getElementById('startingTypeSelect');
firstDropDown.addEventListener('change', async function() { generateDestinationTypeSelect();});

