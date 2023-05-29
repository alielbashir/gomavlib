import pandas as pd
import matplotlib.pyplot as plt

# Increase the font size
plt.rcParams['font.size'] = 16

# Read the CSV data
data = pd.read_csv('results.csv')

# Convert 'Time' to datetime objects
data['Time'] = pd.to_datetime(data['Time'])

# Adjust 'Time' to be seconds from the start
start_time = data['Time'].iloc[0]
data['Time'] = (data['Time'] - start_time).dt.total_seconds()

# Calculate the moving averages with a window size of 5
window_size = 5
data['Channel A Loss MA'] = data['Channel A Loss (%)'].rolling(window=window_size).mean()
data['Channel B Loss MA'] = data['Channel B Loss (%)'].rolling(window=window_size).mean()
data['Filtered Loss MA'] = data['Filtered Loss (%)'].rolling(window=window_size).mean()

# Replace values below 0 with 0
data['Channel A Loss MA'] = data['Channel A Loss MA'].clip(lower=0)
data['Channel B Loss MA'] = data['Channel B Loss MA'].clip(lower=0)
data['Filtered Loss MA'] = data['Filtered Loss MA'].clip(lower=0)

# Remove data after 224 seconds
data = data[data['Time'] <= 224]

# Create a figure and axes with larger dimensions
fig, ax = plt.subplots(figsize=(12, 8))

# Plot the moving averages with increased linewidth
ax.plot(data['Time'].values, data['Channel A Loss MA'].values, label='Bağlantı 1', linewidth=2)
ax.plot(data['Time'].values, data['Channel B Loss MA'].values, label='Bağlantı 2', linewidth=2)
ax.plot(data['Time'].values, data['Filtered Loss MA'].values, label='Tekilleştirilmiş', linewidth=2)

# Add labels to the axes with increased font size
ax.set_xlabel('Zaman (Saniye)', fontsize=20)
ax.set_ylabel('Paket Kaybı (%)', fontsize=20)

# Add a legend with increased font size
ax.legend(loc='upper left', fontsize=18)

# Add gridlines
ax.grid(True)

# Remove excess whitespace
fig.tight_layout()

# Show the plot
plt.show()
