#!/usr/bin/env python3

__author__ = 'Shane Young'
__version__ = '1.0'
__email__ = 'thedzy@thedzy.com'
__date__ = '2023-08-18'
__credits__ = ''

__description__ = \
    """
    write_json_shapes.py: 
    for title_creator
    """

import argparse
import logging
from logging.handlers import RotatingFileHandler
import pprint
from pathlib import Path
import numpy


class FormatPrinter(pprint.PrettyPrinter):

    def __init__(self, width=500, indent=1, depth=None, stream=None, formats={}):
        super().__init__(width=width, indent=indent, depth=depth, stream=stream)
        self.formats = formats

    def format(self, obj, ctx, maxlvl, lvl):
        if type(obj) in self.formats:
            return self.formats[type(obj)] % obj, 1, 0
        return pprint.PrettyPrinter.format(self, obj, ctx, maxlvl, lvl)


class ColourFormat(logging.Formatter):
    """
    Add colour to logging events
    """

    def __init__(self, fmt: str = None, datefmt: str = None, style: str = '%', levels={}) -> None:
        """
        Initialise the formatter
        ft: (str) Format String
        datefmt: (str) Date format
        style: (str) Format style
        levels: tuple, tuple (level number start, colour, attribute
        """
        self.levels = {}
        set_levels = {10: 90, 20: 92, 30: 93, 40: 91, 50: (41, 97)}
        set_levels.update(levels)

        for key in sorted(set_levels.keys()):
            value = set_levels[key]
            colour = str(value) if isinstance(value, (str, int)) else ';'.join(map(str, value))

            self.levels[key] = f'\x1b[5;{colour};m'

        super().__init__(fmt, datefmt, style)

    def formatMessage(self, record: logging.LogRecord, **kwargs: dict) -> str:
        """
        Override the formatMessage method to add colour
        """
        no_colour = u'\x1b[0m'
        for level in self.levels:
            colour = self.levels[level] if record.levelno >= level else colour

        return f'{colour}{super().formatMessage(record, **kwargs)}{no_colour}'


def main() -> None:
    logger.note('Start')

    # Set the dimensions
    size = options.resolution

    # Start a dictionary
    lines = {}

    # Create the shapes for - and |
    logger.info(f'Render -: {options.hyphen}')
    logger.info(f'Render |: {options.pipe}')
    char_increment = 0
    steps = 1
    for line in range(0, size + steps, steps):
        for width in range(line + steps, size + steps, steps):
            char_increment += 1

            empty_matrix1 = numpy.full((size, size), 0)
            empty_matrix2 = numpy.full((size, size), 0)
            empty_matrix1[line:width, :] = 255  # Set the row to 255
            empty_matrix2[:, line:width] = 255  # Set the row to 255

            if not numpy.sum(empty_matrix1) == 0 and not numpy.sum(empty_matrix1) == 255 * size * size:
                if options.hyphen:
                    lines[f'-{char_increment:04d}'] = empty_matrix1.tolist()
            if not numpy.sum(empty_matrix2) == 0 and not numpy.sum(empty_matrix2) == 255 * size * size:
                if options.pipe:
                    lines[f'|{char_increment:04d}'] = empty_matrix2.tolist()


    # Create the shapes for / and \
    logger.info(f'Render /: {options.forward_slash}')
    logger.info(f'Render \: {options.back_slash}')
    for shape in range(0, size):
        if shape % 2 == 0 and shape < size - int(size/4):
            empty_matrix1 = numpy.full((size, size), 255)
            empty_matrix2 = numpy.full((size, size), 255)
            empty_matrix3 = numpy.full((size, size), 255)
            empty_matrix4 = numpy.full((size, size), 255)
            for row in range(0, shape):
                empty_matrix1[size-1 - row, :] = 0  # Set the row to 255
                empty_matrix2[size-1 - row, :] = 0  # Set the row to 255
                empty_matrix3[row, :] = 0  # Set the row to 255
                empty_matrix4[row, :] = 0  # Set the row to 255
                pass
            for row in range(shape, size):
                empty_matrix1[size-1 - row, row - shape:size] = 0
                empty_matrix2[size-1 - row, 0:size - row + shape] = 0
                empty_matrix3[row, row - shape:size] = 0
                empty_matrix4[row, 0:size - row + shape] = 0  # Set the row to 255
                pass
            if options.forward_slash:
                lines[f'/{shape:04d}'] = empty_matrix1.tolist()
                lines[f'/{shape + size:04d}'] = empty_matrix4.tolist()
            if options.back_slash:
                lines[f'\\\\{shape:04d}'] = empty_matrix2.tolist()
                lines[f'\\\\{shape + size:04d}'] = empty_matrix3.tolist()

    # Create the shapes for +
    logger.info(f'Render +: {options.plus}')
    if options.plus:
        mid = int(size / 2)
        empty_matrix1 = numpy.full((size, size), 0)
        empty_matrix1[0:mid, 0:mid] = 255
        lines[f'+{0:04d}'] = empty_matrix1.tolist()
        empty_matrix2 = numpy.full((size, size), 0)
        empty_matrix2[mid:size, 0:mid] = 255
        lines[f'+{1:04d}'] = empty_matrix2.tolist()
        empty_matrix3 = numpy.full((size, size), 0)
        empty_matrix3[mid:size, mid:size] = 255
        lines[f'+{2:04d}'] = empty_matrix3.tolist()
        empty_matrix4 = numpy.full((size, size), 0)
        empty_matrix4[0:mid, mid:size] = 255
        lines[f'+{3:04d}'] = empty_matrix4.tolist()

    # Create the shapes for a filled character
    faces_border = 0 # Add a border to the top/bottom
    side_border = 0 # Add a border to the left/right
    empty_matrix1 = numpy.full((size, size), 0)
    empty_matrix1[faces_border:size - faces_border, side_border:size - side_border] = 255

    # Create the shapes for empty character
    faces_border = 0 # Add a border to the top/bottom
    side_border = 0 # Add a border to the left/right
    empty_matrix2 = numpy.full((size, size), 255)
    empty_matrix2[faces_border:size - faces_border, side_border:size - side_border] = 0

    lines[f' {0:04d}'] = empty_matrix1.tolist()
    lines[f' {1:04d}'] = empty_matrix2.tolist()

    # Output to screen and file
    pp = FormatPrinter(formats={float: "%.2f", int: "%3d", str: "\"%s\""}, width=int(size*5)+15)
    if options.print:
        pp.pprint(lines)
    logger.info(f'Shapes: {len(lines)}')

    with open(options.file, 'w') as f:
        f.write(pp.pformat(lines))

    logger.note('Done')


def create_logger(name: str = __file__, levels: dict = {}) -> logging.Logger:
    """
    Create a logger
    :param name: (str) Name of logger
    :param levels: (dict) Custom log levels
    :return: (logging.Logger) Logger
    """

    # Create log level
    def make_log_level(level_name: str, level_int: int) -> None:
        logging.addLevelName(level_int, level_name.upper())
        setattr(new_logger, level_name, lambda *args: new_logger.log(level_int, *args))

    # Setup logging
    new_logger = logging.getLogger(name)
    new_logger.setLevel(options.debug)

    # Create stream handler
    log_stream_handle = logging.StreamHandler()
    log_format = '[{asctime}] [{levelname:8}] {message}'
    log_stream_handle.setFormatter(ColourFormat('{message}', style='{', levels={20: 16, 21: 92}))
    new_logger.addHandler(log_stream_handle)

    # Set file handler
    if options.output:
        log_size_mb = 6
        log_size = log_size_mb * 1024 * 1024
        log_file_handle = RotatingFileHandler(options.output,
                                              maxBytes=log_size,
                                              backupCount=1
                                              )
        log_file_handle.setFormatter(logging.Formatter(log_format, style='{'))
        new_logger.addHandler(log_file_handle)

    # Create custom levels
    for level in levels.items():
        make_log_level(*level)

    return new_logger


if __name__ == '__main__':
    def valid_path(path):
        parent = Path(path).parent
        if not parent.is_dir():
            print(f'{parent} is not a directory, make it?')
            if input('y/n: ').lower() == 'y':
                parent.mkdir(parents=True, exist_ok=True)
                Path(path)
            raise argparse.ArgumentTypeError(f'{path} is not a directory')
        return Path(path)


    # Create argument parser
    parser = argparse.ArgumentParser(description=__description__)

    parser.add_argument('-r', '--resolution', type=int, default=8,
                        action='store', dest='resolution',
                        help='resolution of the json')

    parser.add_argument('-f', '--file', type=valid_path,
                        action='store', dest='file',
                        required=True,
                        help='output file')

    parser.add_argument('-p', '--no-print', default=True,
                        action='store_false', dest='print',
                        help='no dont print to screen')

    omits = parser.add_argument_group('Omit shapes')
    omits.add_argument('--pipe', default=True,
                        action='store_false', dest='pipe',
                        help='omit pipe')

    omits.add_argument('--hyphen', default=True,
                        action='store_false', dest='hyphen',
                        help='omit hyphen')

    omits.add_argument('--forward-slash', default=True,
                        action='store_false', dest='forward_slash',
                        help='omit forward-slash')

    omits.add_argument('--back-slash', default=True,
                        action='store_false', dest='back_slash',
                        help='omit back-slash')

    omits.add_argument('--plus', default=True,
                        action='store_false', dest='plus',
                        help='omit plus')

    # Debug option
    parser.add_argument('--debug', default=20,
                        action='store_const', dest='debug', const=10,
                        help=argparse.SUPPRESS)

    # Output
    parser.add_argument('-o', '--output', type=valid_path,
                        default=Path('/tmp').joinpath(Path(__file__).stem).with_suffix('.log'),
                        action='store', dest='output',
                        help='output log')

    options = parser.parse_args()

    logger = create_logger(levels={'note': 21})
    logger.debug('Debug ON')
    logger.debug(pprint.pformat(options))

    main()
