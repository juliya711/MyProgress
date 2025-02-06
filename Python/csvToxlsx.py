import pandas as pd

# csv_file = "C:\\Users\\juliyapraisy.j\\Downloads\\10-09-2024-11-15-11-am- Vulnerablity Details.csv"

excel_file = "output.xlsx"

df = pd.read_csv(csv_file)

df.to_excel(excel_file, index=False, engine='openpyxl')

print(f"CSV successfully converted to Excel: {excel_file}")