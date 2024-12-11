from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.common.action_chains import ActionChains
from selenium.common.exceptions import NoSuchElementException
from webdriver_manager.chrome import ChromeDriverManager
from textblob import TextBlob
import matplotlib.pyplot as plt
import time

# Set up the Selenium WebDriver
driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()))

# Open the Amazon review page
# url = 'https://www.amazon.in/Acer-15-6-inch-Graphics-Obsidian-AN515-54/dp/B088FLW4TW/ref=cm_cr_srp_d_product_top?ie=UTF8&th=1%27'
url='https://www.flipkart.com/acer-nitro-5-core-i5-9th-gen-8-gb-1-tb-hdd-windows-10-home-3-gb-graphics-nvidia-geforce-gtx-1050-an515-54-563y-an515-54-52h2-gaming-laptop/product-reviews/itmcc400f97b66f1?pid=COMFHNY8HYPHY7ZH&lid=LSTCOMFHNY8HYPHY7ZHIIO4JI&sortOrder=POSITIVE_FIRST&certifiedBuyer=false&aid=overall'
# url = 'https://www.amazon.in/dp/B088FLW4TW?ref=cm_sw_r_cso_wa_apin_dp_CFEXDADWYFMX81GXDZJW&ref_=cm_sw_r_cso_wa_apin_dp_CFEXDADWYFMX81GXDZJW&social_share=cm_sw_r_cso_wa_apin_dp_CFEXDADWYFMX81GXDZJW&starsLeft=1&skipTwisterOG=1'
driver.get(url)

# Wait for the page to load fully
time.sleep(5)

# Define common terms for specific features
common_terms = ['battery', 'camera', 'trackpad', 'screen quality', 'battery life', 'heat', 'speaker', 'graphic', 'keyboard']

# Dictionaries to store keyword counts for common terms under sentiment categories
term_sentiment_count = {term: {'positive': 0, 'negative': 0, 'neutral': 0} for term in common_terms}

# Lists to store sentiment polarities for each review
polarities = []

# Pagination loop
while True:
    try:
        # Wait and dynamically locate review elements
        time.sleep(3)  # Adjust the sleep time if necessary
        reviews = driver.find_elements(By.XPATH, '//span[@data-hook="review-body"]')

        # Analyze the polarity of each review
        for review in reviews:
            review_text = review.text.lower()

            # Get the polarity score using TextBlob
            blob = TextBlob(review_text)
            polarity = blob.sentiment.polarity  # Polarity score ranges from -1 to 1

            # Append the polarity to the list
            polarities.append(polarity)

            # Count specific terms under sentiment categories
            for term in common_terms:
                if term in review_text:
                    # Determine sentiment based on polarity score
                    if polarity > 0:
                        term_sentiment_count[term]['positive'] += 1
                    elif polarity < 0:
                        term_sentiment_count[term]['negative'] += 1
                    else:
                        term_sentiment_count[term]['neutral'] += 1

        # Locate and click the "Next Page" button
        next_button = driver.find_element(By.XPATH, '//li[@class="a-last"]/a')
        ActionChains(driver).move_to_element(next_button).click(next_button).perform()
    except NoSuchElementException:
        # If "Next Page" button is not found, break the loop
        print("No more pages to scrape.")
        break

# Prepare data for overall sentiment count
positive_count = sum(1 for p in polarities if p > 0)
negative_count = sum(1 for p in polarities if p < 0)
neutral_count = sum(1 for p in polarities if p == 0)

overall_sentiment = ['Positive', 'Negative', 'Neutral']
overall_values = [positive_count, negative_count, neutral_count]

# Prepare data for the keyword sentiment breakdown (for common terms)
x_pos = range(len(common_terms))
positive_values = [term_sentiment_count[term]['positive'] for term in common_terms]
negative_values = [term_sentiment_count[term]['negative'] for term in common_terms]
neutral_values = [term_sentiment_count[term]['neutral'] for term in common_terms]

# Create the figure and subplots for both charts
fig, axes = plt.subplots(1, 2, figsize=(15, 6))

# Chart 1: Overall sentiment
axes[0].bar(overall_sentiment, overall_values, color=['green', 'red', 'yellow'])
axes[0].set_title('Overall Sentiment Analysis')
axes[0].set_xlabel('Sentiment Categories')
axes[0].set_ylabel('Count')

# Chart 2: Keyword sentiment breakdown (multi-bar chart for common terms)
width = 0.25  # Bar width for multi-bar chart
axes[1].bar([x - width for x in x_pos], positive_values, width, label='Positive', color='green')
axes[1].bar(x_pos, negative_values, width, label='Negative', color='red')
axes[1].bar([x + width for x in x_pos], neutral_values, width, label='Neutral', color='yellow')

axes[1].set_title('Keyword Sentiment Analysis (Common Terms)')
axes[1].set_xlabel('Common Terms')
axes[1].set_ylabel('Count')
axes[1].set_xticks(x_pos)
axes[1].set_xticklabels(common_terms, rotation=45, ha='right')
axes[1].legend()

# Display the charts
plt.tight_layout()
plt.show()

# Close the browser
driver.quit()