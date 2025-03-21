import { PRODUCTS } from "@/assets/data";
import ProductCard from "@/libs/globals/components/cards/ProductCard";
import MainLayout from "@/libs/globals/components/layouts/MainLayout";
import SectionTitle from "@/libs/globals/components/sections/SectionTitle";
import CollectionsSection from "@/libs/routes/home/CollectionsSection";
import FeatureProductsGallery from "@/libs/routes/home/FeatureProductsGallery";
import InstagramFeed from "@/libs/routes/home/InstagramFeed";
import MainHeroGallery from "@/libs/routes/home/MainHeroGallery";
import { createFileRoute, Link } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  head: () => ({
    meta: [{ title: "Home Page" }],
  }),
  component: HomePage,
});

function HomePage() {
  return (
    <MainLayout>
      <div className="flex w-full flex-col gap-20 p-6 lg:p-[auto]">
        <section className="w-full">
          <MainHeroGallery />
        </section>
        <section className="flex w-full flex-col items-start md:gap-3 xl:gap-6 2xl:flex-row">
          <SectionTitle className="mb-3 md:mr-[inherit] xl:mr-56">
            Products
          </SectionTitle>
          <div className="flex w-full flex-1 flex-col flex-wrap items-stretch gap-6 sm:flex-row lg:gap-6">
            {PRODUCTS.slice(0, 6).map((p, i) => (
              <Link
                key={`product-card-${i}`}
                to="/products/$productId"
                params={{ productId: p.id.toString() }}
              >
                <ProductCard
                  title={p.name}
                  variants={p.colorOptions.map((v) => ({
                    imageUrl: v.imageUrl,
                    colorSwatch: v.colorSwatch,
                    price: 100.0 * Math.random(),
                  }))}
                  onClick={(e) => {
                    e.preventDefault();
                  }}
                />
              </Link>
            ))}
          </div>
        </section>
        <section>
          <SectionTitle>Categories</SectionTitle>
          <CollectionsSection />
        </section>
        <section>
          <SectionTitle>Clearance Sale</SectionTitle>
          <FeatureProductsGallery />
        </section>
        <section>
          <SectionTitle>Our Instagram</SectionTitle>
          <InstagramFeed />
        </section>
      </div>
    </MainLayout>
  );
}
