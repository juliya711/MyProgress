import pandas as pd
import sys
from io import StringIO

input_stream = sys.stdin.read()

df = pd.read_csv(StringIO(input_stream))

import io
output = io.BytesIO()
df.to_excel(output, index=False, engine='openpyxl')

output.seek(0)

sys.stdout.buffer.write(output.read())