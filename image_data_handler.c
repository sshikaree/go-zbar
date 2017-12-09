// #include <stdio.h>
#include <zbar.h>
#include "image_data_handler.h"

extern void image_handler_callback(zbar_image_t *image, const void *userdata);


void image_data_handler(zbar_image_t *image, const void *userdata) {
    image_handler_callback(image, userdata);
}


// void image_data_handler(zbar_image_t *image, const void *userdata) {
//     /* extract results */
//     const zbar_symbol_t *symbol = zbar_image_first_symbol(image);
//     for(; symbol; symbol = zbar_symbol_next(symbol)) {
//         /* do something useful with results */
//         zbar_symbol_type_t typ = zbar_symbol_get_type(symbol);
//         const char *data = zbar_symbol_get_data(symbol);
//         printf("decoded %s symbol \"%s\"\n", zbar_get_symbol_name(typ), data);
//     }
// }