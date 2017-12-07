package zbar

// "color" of element: bar or space.
const (
	ZBAR_SPACE = iota // light area or space between bars
	ZBAR_BAR          // dark area or colored bar segment
)

// zbar_symbol_type_t
// decoded symbol type.
const (
	ZBAR_NONE        = 0   /**< no symbol decoded */
	ZBAR_PARTIAL     = 1   /**< intermediate status */
	ZBAR_EAN2        = 2   /**< GS1 2-digit add-on */
	ZBAR_EAN5        = 5   /**< GS1 5-digit add-on */
	ZBAR_EAN8        = 8   /**< EAN-8 */
	ZBAR_UPCE        = 9   /**< UPC-E */
	ZBAR_ISBN10      = 10  /**< ISBN-10 (from EAN-13). @since 0.4 */
	ZBAR_UPCA        = 12  /**< UPC-A */
	ZBAR_EAN13       = 13  /**< EAN-13 */
	ZBAR_ISBN13      = 14  /**< ISBN-13 (from EAN-13). @since 0.4 */
	ZBAR_COMPOSITE   = 15  /**< EAN/UPC composite */
	ZBAR_I25         = 25  /**< Interleaved 2 of 5. @since 0.4 */
	ZBAR_DATABAR     = 34  /**< GS1 DataBar (RSS). @since 0.11 */
	ZBAR_DATABAR_EXP = 35  /**< GS1 DataBar Expanded. @since 0.11 */
	ZBAR_CODABAR     = 38  /**< Codabar. @since 0.11 */
	ZBAR_CODE39      = 39  /**< Code 39. @since 0.4 */
	ZBAR_PDF417      = 57  /**< PDF417. @since 0.6 */
	ZBAR_QRCODE      = 64  /**< QR Code. @since 0.10 */
	ZBAR_CODE93      = 93  /**< Code 93. @since 0.11 */
	ZBAR_CODE128     = 128 /**< Code 128 */
)

// ZBar orientation
const (
	ZBAR_ORIENT_UNKNOWN = -1 + iota /**< unable to determine orientation */
	ZBAR_ORIENT_UP                  /**< upright, read left to right */
	ZBAR_ORIENT_RIGHT               /**< sideways, read top to bottom */
	ZBAR_ORIENT_DOWN                /**< upside-down, read right to left */
	ZBAR_ORIENT_LEFT                /**< sideways, read bottom to top */
)

// zbar_config_t
// decoder configuration options
const (
	ZBAR_CFG_ENABLE     = iota // enable symbology/feature
	ZBAR_CFG_ADD_CHECK         // enable check digit when optional
	ZBAR_CFG_EMIT_CHECK        // return check digit when present
	ZBAR_CFG_ASCII             // enable full ASCII character set
	ZBAR_CFG_NUM               // number of boolean decoder configs
	ZBAR_CFG_MIN_LEN           // minimum data length for valid decode
	ZBAR_CFG_MAX_LEN           // maximum data length for valid decode
	ZBAR_CFG_POSITION          // enable scanner to collect position data
	ZBAR_CFG_X_DENSITY         // image scanner vertical scan density
	ZBAR_CFG_Y_DENSITY         // image scanner horizontal scan density
)

// error codes
const (
	ZBAR_OK              = iota // no error
	ZBAR_ERR_NOMEM              // out of memory
	ZBAR_ERR_INTERNAL           // internal library error
	ZBAR_ERR_UNSUPPORTED        // unsupported request
	ZBAR_ERR_INVALID            // invalid request
	ZBAR_ERR_SYSTEM             // system error
	ZBAR_ERR_LOCKING            // locking error
	ZBAR_ERR_BUSY               // all resources busy
	ZBAR_ERR_XDISPLAY           // X11 display error.
	ZBAR_ERR_XPROTO             // X11 protocol error.
	ZBAR_ERR_CLOSED             // output window is closed
	ZBAR_ERR_WINAPI             // windows system error
	ZBAR_ERR_NUM                // number of error codes
)

// Decoder symbology modifier flags
const (
	ZBAR_MOD_GS1 = iota // barcode tagged as GS1 (EAN.UCC) reserved (eg, FNC1 before first data character).data may be parsed as a sequence of GS1 AIs
	ZBAR_MOD_AIM        // barcode tagged as AIM reserved (eg, FNC1 after first character or digit pair)
	ZBAR_MOD_NUM        // number of modifiers
)
