#include <zbar.h>
#include <string.h>
// #include <stdio.h>
//#include <stdlib.h>

int scan_single_symbol(const char *device, char *result, char *code_type) {
	/* create a Processor */
	zbar_processor_t *proc = zbar_processor_create(1);
	/* configure the Processor */
	zbar_processor_set_config(proc, 0, ZBAR_CFG_ENABLE, 1);
	/* initialize the Processor */
	zbar_processor_init(proc, device, 1);

	/* enable the preview window */
	zbar_processor_set_visible(proc, 1);
	zbar_processor_set_active(proc, 1);

	int ok = zbar_process_one(proc, -1);
	if (ok < 0) {
		//fprintf(stderr, "Processing error.");
		zbar_processor_destroy(proc);
		return -1;
	}
	const zbar_symbol_set_t *results = zbar_processor_get_results(proc);
	const zbar_symbol_t *symbol = zbar_symbol_set_first_symbol(results);
	zbar_symbol_type_t typ = zbar_symbol_get_type(symbol);

	strcpy(result, (char*)zbar_symbol_get_data(symbol));
	strcpy(code_type, (char*)zbar_get_symbol_name(typ));

	/* clean up */
	zbar_processor_destroy(proc);
	return 0;
}


