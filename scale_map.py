#!/usr/bin/env python3

__author__ = 'Shane Young'
__version__ = '1.0'
__email__ = 'thedzy@thedzy.com'
__date__ = '2023-08-25'
__credits__ = ''

__description__ = \
    """
    scale_map.py: 
    Scale up or down a character map
    """

import argparse
import logging
from logging.handlers import RotatingFileHandler
import pprint
from pathlib import Path
import numpy
import json
from scipy.ndimage import zoom




class FormatPrinter(pprint.PrettyPrinter):
    """
    Override the formatt method to format by specification
    """
    def __init__(self, width=500, indent=1, depth=None, stream=None, formats={}):
        super().__init__(width=width, indent=indent, depth=depth, stream=stream)
        self.formats = formats

    def format(self, obj, ctx, maxlvl, lvl):
        if type(obj) in self.formats:
            return self.formats[type(obj)] % obj, 1, 0
        return pprint.PrettyPrinter.format(self, obj, ctx, maxlvl, lvl)


def main() -> None:

    # Load the input file
    data = json.load(options.input_file)

    # Get old and new sizes
    new_size = options.new_size
    for _, matrix_value in data.items():
        old_size = len(matrix_value)
        break
    if new_size == 0:
        new_size = old_size * 2

    # print the old map
    if options.print:
        pp = FormatPrinter(formats={float: "%.2f", int: "%3d", str: "\"%s\""}, width=int(old_size*5)+15)
        print('Original')
        pp.pprint(data)

    # Scale up each matrix with the scipy zoom
    for key, matrix_value in data.items():
        # Convert the matrix to a numpy array
        matrix_numpy = numpy.array(matrix_value)

        # Scale the matrix using scipy's zoom function
        resized_matrix = zoom(matrix_numpy, new_size / old_size, order=options.mode)

        # Update the value in the JSON object
        data[key] = resized_matrix.tolist()

    # Output to screen and file
    pp = FormatPrinter(formats={float: "%.2f", int: "%3d", str: "\"%s\""}, width=int(new_size*5)+15)
    if options.print:
        print('Resized')
        pp.pprint(data)
        print(f'Shapes: {len(data)}')

    with open(options.output_file, 'w') as f:
        f.write(pp.pformat(data))


    print(f"{len(data)} characters have been resized from {old_size}x{old_size} to {new_size}x{new_size} and saved as {options.output_file}")


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

    def int_range(minimum, maximum):
        def int_range_(value):
            int_value = int(value)
            if minimum <= int_value <= maximum:
                return int_value
            else:
                raise argparse.ArgumentTypeError(f'{value} is not in range {minimum}-{maximum}')
        return int_range_


    # Create argument parser
    parser = argparse.ArgumentParser(description=__description__)

    parser.add_argument('-n', '--new_size', type=int, default=0,
                        action='store', dest='new_size',
                        required=False,
                        help='new size (default double)')

    parser.add_argument('-m', '--mode', type=int_range(0,5), default=0,
                        action='store', dest='mode',
                        required=False,
                        help='resizing mode (0-5)')

    parser.add_argument('-i', '--input', type=argparse.FileType('r'),
                        action='store', dest='input_file',
                        required=True,
                        help='input file')

    parser.add_argument('-o', '--output', type=valid_path,
                        action='store', dest='output_file',
                        required=True,
                        help='output file')

    parser.add_argument('-p', '--print', default=False,
                        action='store_true', dest='print',
                        help='print output')

    # Debug option
    parser.add_argument('--debug', default=20,
                        action='store_const', dest='debug', const=10,
                        help=argparse.SUPPRESS)

    options = parser.parse_args()

    main()
