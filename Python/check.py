from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from webdriver_manager.chrome import ChromeDriverManager
from textblob import TextBlob  # Import TextBlob for polarity analysis
import matplotlib.pyplot as plt
import time

# Set up the Selenium WebDriver
driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()))

# Open the Amazon review page
url = 'https://www.amazon.in/dp/B0CM5JLSNS/ref=sspa_dk_detail_1?psc=1&pd_rd_i=B0CM5JLSNS&pd_rd_w=J364W&content-id=amzn1.sym.9f1cb690-f0b7-44de-b6ff-1bad1e37d3f0&pf_rd_p=9f1cb690-f0b7-44de-b6ff-1bad1e37d3f0&pf_rd_r=Y41QBBF8P245ZG7YF7GZ&pd_rd_wg=93zMC&pd_rd_r=0c9c89d7-bc89-47f7-8d83-d1fe1d21fcf7&sp_csd=d2lkZ2V0TmFtZT1zcF9kZXRhaWxfdGhlbWF0aWM'
driver.get(url)

# Wait for the page to load fully
time.sleep(5)  # Adjust the sleep time if necessary

# Dynamically locate review elements
reviews = driver.find_elements(By.XPATH, '//span[@data-hook="review-body"]')

# List of common terms for analysis
common_terms = ['battery', 'camera', 'trackpad', 'screen quality', 'battery life', 'heat', 'speaker', 'graphic', 'keyboard']

# Dictionaries to store keyword counts for sentiment categories
term_sentiment_count = {term: {'positive': 0, 'negative': 0, 'neutral': 0} for term in common_terms}

# Lists to store sentiment polarities for each review
polarities = []

# Analyze the sentiment and common term occurrences for each review
for review in reviews:
    review_text = review.text.lower()
    
    # Get the polarity score using TextBlob
    blob = TextBlob(review_text)
    polarity = blob.sentiment.polarity  # Polarity score ranges from -1 to 1
    
    # Append the polarity to the list
    polarities.append(polarity)
    
    # Check if common terms are mentioned and associate them with sentiment
    for term in common_terms:
        if term in review_text:
            if polarity > 0:
                term_sentiment_count[term]['positive'] += 1
            elif polarity < 0:
                term_sentiment_count[term]['negative'] += 1
            else:
                term_sentiment_count[term]['neutral'] += 1

# Prepare data for overall sentiment count
positive_count = sum(1 for p in polarities if p > 0)
negative_count = sum(1 for p in polarities if p < 0)
neutral_count = sum(1 for p in polarities if p == 0)

overall_sentiment = ['Positive', 'Negative', 'Neutral']
overall_values = [positive_count, negative_count, neutral_count]

# Prepare data for the keyword sentiment breakdown chart
x_pos = range(len(common_terms))
positive_values = [term_sentiment_count[term]['positive'] for term in common_terms]
negative_values = [term_sentiment_count[term]['negative'] for term in common_terms]
neutral_values = [term_sentiment_count[term]['neutral'] for term in common_terms]

# Create the figure and subplots for overall sentiment and keyword sentiment breakdown
fig, axes = plt.subplots(1, 2, figsize=(15, 6))

# Chart 1: Overall sentiment
axes[0].bar(overall_sentiment, overall_values, color=['green', 'red', 'yellow'])
axes[0].set_title('Overall Sentiment Analysis')
axes[0].set_xlabel('Sentiment Categories')
axes[0].set_ylabel('Count')

# Chart 2: Keyword sentiment breakdown (multi-bar chart)
width = 0.25  # Bar width
axes[1].bar([x - width for x in x_pos], positive_values, width, label='Positive', color='green')
axes[1].bar(x_pos, negative_values, width, label='Negative', color='red')
axes[1].bar([x + width for x in x_pos], neutral_values, width, label='Neutral', color='yellow')

axes[1].set_title('Keyword Sentiment Analysis')
axes[1].set_xlabel('Keywords')
axes[1].set_ylabel('Count')
axes[1].set_xticks(x_pos)
axes[1].set_xticklabels(common_terms, rotation=45, ha='right')
axes[1].legend()

# Display the charts
plt.tight_layout()
plt.show()

# Close the browser
driver.quit()